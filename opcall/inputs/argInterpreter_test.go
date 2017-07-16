package inputs

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"path/filepath"
	"strconv"
)

var _ = Context("argInterpreter", func() {
	Context("Interpret", func() {
		Context("param nil", func() {
			It("should return expected error", func() {
				/* arrange */
				providedName := "dummyName"

				expectedError := fmt.Errorf("Unable to bind to '%v'; '%v' not a defined input", providedName, providedName)

				objectUnderTest := _argInterpreter{}

				/* act */
				_, actualError := objectUnderTest.Interpret(
					"dummyName",
					"dummyValue",
					nil,
					"dummyPkgPath",
					map[string]*model.Value{},
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("Implicit arg", func() {
			Context("Ref not in scope", func() {
				It("should return expected error", func() {
					/* arrange */
					providedName := "dummyName"

					expectedError := fmt.Errorf("Unable to bind to '%v' via implicit ref; '%v' not in scope", providedName, providedName)

					objectUnderTest := _argInterpreter{}

					/* act */
					_, actualError := objectUnderTest.Interpret(
						providedName,
						"",
						&model.Param{},
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualError).To(Equal(expectedError))
				})
			})
			Context("Ref in scope", func() {
				It("should call validate.Validate w/ expected args & return result", func() {
					/* arrange */
					providedName := "dummyName"
					providedValue := ""
					providedParam := &model.Param{}
					expectedValue := &model.Value{String: new(string)}
					providedScope := map[string]*model.Value{providedName: expectedValue}

					objectUnderTest := _argInterpreter{}

					/* act */
					actualValue, actualError := objectUnderTest.Interpret(
						providedName,
						providedValue,
						providedParam,
						"dummyPkgPath",
						providedScope,
					)

					/* assert */
					Expect(actualValue).To(Equal(expectedValue))
					Expect(actualError).To(BeNil())
				})
			})
		})
		Context("Deprecated explicit arg", func() {
			It("should call validate.Validate w/ expected args & return result", func() {
				/* arrange */
				providedValue := "dummyValue"
				providedParam := &model.Param{}
				expectedValue := &model.Value{String: new(string)}
				providedScope := map[string]*model.Value{providedValue: expectedValue}

				objectUnderTest := _argInterpreter{}

				/* act */
				actualValue, actualError := objectUnderTest.Interpret(
					"dummyName",
					providedValue,
					providedParam,
					"dummyPkgPath",
					providedScope,
				)

				/* assert */
				Expect(actualValue).To(Equal(expectedValue))
				Expect(actualError).To(BeNil())
			})
		})
		Context("Explicit arg", func() {
			Context("Ref not in scope", func() {
				It("should return expected error", func() {
					/* arrange */
					providedName := "dummyName"
					explicitRef := "dummyRef"
					providedValue := fmt.Sprintf("$(%v)", explicitRef)
					providedParam := &model.Param{}

					expectedError := fmt.Errorf("Unable to bind '%v' to '%v' via explicit ref; '%v' not in scope", providedName, explicitRef, explicitRef)

					objectUnderTest := _argInterpreter{}

					/* act */
					_, actualError := objectUnderTest.Interpret(
						providedName,
						providedValue,
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualError).To(Equal(expectedError))
				})
			})
			Context("Ref in scope", func() {
				It("should call validate.Validate w/ expected args & return result", func() {
					/* arrange */
					explicitRef := "dummyRef"
					providedValue := fmt.Sprintf("$(%v)", explicitRef)
					providedParam := &model.Param{}
					expectedValue := &model.Value{String: new(string)}
					providedScope := map[string]*model.Value{explicitRef: expectedValue}

					objectUnderTest := _argInterpreter{}

					/* act */
					actualValue, actualError := objectUnderTest.Interpret(
						"dummyName",
						providedValue,
						providedParam,
						"dummyPkgPath",
						providedScope,
					)

					/* assert */
					Expect(actualValue).To(Equal(expectedValue))
					Expect(actualError).To(BeNil())
				})
			})
		})
		Context("Interpolated arg", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedValue := "dummyValue"
				providedParam := &model.Param{}
				providedScope := map[string]*model.Value{"dummyScopeRef": {}}

				fakeInterpolater := new(interpolater.Fake)

				objectUnderTest := _argInterpreter{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.Interpret(
					"dummyName",
					providedValue,
					providedParam,
					"dummyPkgPath",
					providedScope,
				)

				/* assert */
				actualTemplate, actualScope := fakeInterpolater.InterpolateArgsForCall(0)
				Expect(actualTemplate).To(Equal(providedValue))
				Expect(actualScope).To(Equal(providedScope))
			})
			Context("Input is string", func() {
				It("should return expected result", func() {
					/* arrange */
					providedParam := &model.Param{String: &model.StringParam{}}

					fakeInterpolater := new(interpolater.Fake)
					interpolatedValue := "dummyValue"
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedResult := &model.Value{String: &interpolatedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
			})
			Context("Input is Dir", func() {
				Context("bound to pkg dir", func() {
					It("should return expected results", func() {
						/* arrange */
						providedValue := "/somePkgDir"

						providedParam := &model.Param{Dir: &model.DirParam{}}

						providedPkgPath := "/dummyPkgPath"

						fakeInterpolater := new(interpolater.Fake)

						interpolatedValue := "dummyValue"
						fakeInterpolater.InterpolateReturns(interpolatedValue)

						expectedValue := filepath.Join(providedPkgPath, interpolatedValue)

						expectedResult := &model.Value{Dir: &expectedValue}

						objectUnderTest := _argInterpreter{
							interpolater: fakeInterpolater,
						}

						/* act */
						actualResult, actualError := objectUnderTest.Interpret(
							"dummyName",
							providedValue,
							providedParam,
							providedPkgPath,
							map[string]*model.Value{},
						)

						/* assert */
						Expect(actualResult).To(Equal(expectedResult))
						Expect(actualError).To(BeNil())
					})
				})
				It("should return expected result", func() {
					/* arrange */
					providedParam := &model.Param{Dir: &model.DirParam{}}

					fakeInterpolater := new(interpolater.Fake)
					interpolatedValue := "dummyValue"
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedResult := &model.Value{Dir: &interpolatedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
				It("should root path", func() {
					/* arrange */
					providedParam := &model.Param{Dir: &model.DirParam{}}

					fakeInterpolater := new(interpolater.Fake)

					expectedValue := fmt.Sprintf("%v%v", string(filepath.Separator), "dummyValue")
					interpolatedValue := fmt.Sprintf("..\\../%v../..\\", expectedValue)
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedResult := &model.Value{Dir: &expectedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
			})
			Context("Input is Number", func() {
				It("should call validate w/ expected args", func() {
					/* arrange */
					providedParam := &model.Param{Number: &model.NumberParam{}}

					fakeInterpolater := new(interpolater.Fake)
					interpolatedValue := "2.1"
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedValue, err := strconv.ParseFloat(interpolatedValue, 64)
					if nil != err {
						panic(err)
					}
					expectedResult := &model.Value{Number: &expectedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
			})
			Context("Input is File", func() {
				Context("bound to pkg file", func() {
					It("should return expected results", func() {
						/* arrange */
						providedValue := "/somePkgFile"

						providedParam := &model.Param{File: &model.FileParam{}}

						providedPkgPath := "/dummyPkgPath"

						fakeInterpolater := new(interpolater.Fake)

						interpolatedValue := "dummyValue"
						fakeInterpolater.InterpolateReturns(interpolatedValue)

						expectedValue := filepath.Join(providedPkgPath, interpolatedValue)

						expectedResult := &model.Value{File: &expectedValue}

						objectUnderTest := _argInterpreter{
							interpolater: fakeInterpolater,
						}

						/* act */
						actualResult, actualError := objectUnderTest.Interpret(
							"dummyName",
							providedValue,
							providedParam,
							providedPkgPath,
							map[string]*model.Value{},
						)

						/* assert */
						Expect(actualResult).To(Equal(expectedResult))
						Expect(actualError).To(BeNil())
					})
				})
				It("should return expected results", func() {
					/* arrange */
					providedParam := &model.Param{File: &model.FileParam{}}

					fakeInterpolater := new(interpolater.Fake)

					interpolatedValue := "dummyValue"
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedResult := &model.Value{File: &interpolatedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
				It("should root path", func() {
					/* arrange */
					providedParam := &model.Param{File: &model.FileParam{}}

					fakeInterpolater := new(interpolater.Fake)

					expectedValue := fmt.Sprintf("%v%v", string(filepath.Separator), "dummyValue")
					interpolatedValue := fmt.Sprintf("..\\../%v../..\\", expectedValue)
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedResult := &model.Value{File: &expectedValue}

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualResult, actualError := objectUnderTest.Interpret(
						"dummyName",
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualResult).To(Equal(expectedResult))
					Expect(actualError).To(BeNil())
				})
			})
			Context("Input is Socket", func() {
				It("should return expected error", func() {
					/* arrange */
					providedName := "dummyName"
					providedParam := &model.Param{Socket: &model.SocketParam{}}

					fakeInterpolater := new(interpolater.Fake)

					interpolatedValue := "dummyValue"
					fakeInterpolater.InterpolateReturns(interpolatedValue)

					expectedError := fmt.Errorf("Unable to bind '%v' to '%v'; sockets must be passed by reference", providedName, interpolatedValue)

					objectUnderTest := _argInterpreter{
						interpolater: fakeInterpolater,
					}

					/* act */
					_, actualError := objectUnderTest.Interpret(
						providedName,
						"dummyValue",
						providedParam,
						"dummyPkgPath",
						map[string]*model.Value{},
					)

					/* assert */
					Expect(actualError).To(Equal(expectedError))
				})
			})
		})
	})
})