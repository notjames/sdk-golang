package unbracketed

import (
	"errors"
	"fmt"

	"github.com/opspec-io/sdk-golang/op/interpreter/interpolater/dereferencer/identifier/value"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data/coerce"
	"github.com/opspec-io/sdk-golang/model"
)

var _ = Context("DeReferencer", func() {
	Context("NewDeReferencer", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(NewDeReferencer()).Should(Not(BeNil()))
		})
	})
	Context("DeReference", func() {
		It("should call coerce.ToObject w/ expected args", func() {
			/* arrange */
			providedData := model.Value{String: new(string)}

			fakeCoerce := new(coerce.Fake)
			// err to trigger immediate return
			fakeCoerce.ToObjectReturns(nil, errors.New("dummyErr"))

			objectUnderTest := _deReferencer{
				coerce: fakeCoerce,
			}

			/* act */
			objectUnderTest.DeReference(
				"dummyRef",
				&providedData,
			)

			/* assert */
			actualData := fakeCoerce.ToObjectArgsForCall(0)

			Expect(*actualData).To(Equal(providedData))

		})
		Context("coerce.ToObject errs", func() {
			It("should return expected result", func() {

				/* arrange */
				providedRef := "dummyRef"
				providedData := model.Value{String: new(string)}

				fakeCoerce := new(coerce.Fake)
				toObjectErr := errors.New("toObjectErr")
				fakeCoerce.ToObjectReturns(nil, toObjectErr)

				expectedErr := fmt.Errorf("unable to deReference '%v'; error was %v", providedRef, toObjectErr.Error())

				objectUnderTest := _deReferencer{
					coerce: fakeCoerce,
				}

				/* act */
				_, _, actualErr := objectUnderTest.DeReference(
					providedRef,
					&providedData,
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("coerce.ToObject doesn't err", func() {

			It("should call parser.Parse w/ expected args", func() {
				/* arrange */
				providedRef := "dummyRef"

				fakeCoerce := new(coerce.Fake)
				// empty Object to trigger immediate return
				fakeCoerce.ToObjectReturns(&model.Value{Object: map[string]interface{}{}}, nil)

				fakeParser := new(FakeParser)
				fakeParser.ParseReturns("dummyIdentifier", "dummyRefRemainder")

				objectUnderTest := _deReferencer{
					coerce: fakeCoerce,
					parser: fakeParser,
				}

				/* act */
				objectUnderTest.DeReference(
					providedRef,
					nil,
				)

				/* assert */
				actualRef := fakeParser.ParseArgsForCall(0)

				Expect(actualRef).To(Equal(providedRef))

			})
			Context("identifier not in object", func() {
				It("should return expected result", func() {

					/* arrange */
					providedRef := "dummyRef"

					fakeCoerce := new(coerce.Fake)
					fakeCoerce.ToObjectReturns(&model.Value{Object: map[string]interface{}{}}, nil)

					fakeParser := new(FakeParser)
					identifier := "identifier"
					fakeParser.ParseReturns(identifier, "dummyRefRemainder")

					expectedErr := fmt.Errorf("unable to deReference '%v'; '%v' doesn't exist", providedRef, identifier)

					objectUnderTest := _deReferencer{
						coerce: fakeCoerce,
						parser: fakeParser,
					}

					/* act */
					_, _, actualErr := objectUnderTest.DeReference(
						providedRef,
						new(model.Value),
					)

					/* assert */
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
			Context("identifier in object", func() {
				It("should call valueConstructor.Construct w/ expected args", func() {

					/* arrange */
					identifier := "identifier"

					fakeCoerce := new(coerce.Fake)
					objectValue := model.Value{Object: map[string]interface{}{identifier: "dummyValue"}}
					fakeCoerce.ToObjectReturns(&objectValue, nil)

					fakeParser := new(FakeParser)
					fakeParser.ParseReturns(identifier, "dummyRefRemainder")

					fakeValueConstructor := new(value.FakeConstructor)
					fakeValueConstructor.ConstructReturns(nil, errors.New("dummyErr"))

					objectUnderTest := _deReferencer{
						coerce:           fakeCoerce,
						parser:           fakeParser,
						valueConstructor: fakeValueConstructor,
					}

					/* act */
					objectUnderTest.DeReference(
						"dummyRef",
						new(model.Value),
					)

					/* assert */
					actualData := fakeValueConstructor.ConstructArgsForCall(0)
					Expect(actualData).To(Equal(objectValue.Object[identifier]))
				})
				Context("valueConstructor.Construct errs", func() {

					It("should return expected result", func() {

						/* arrange */
						providedRef := "dummyRef"

						identifier := "identifier"

						fakeCoerce := new(coerce.Fake)
						objectValue := model.Value{Object: map[string]interface{}{identifier: "dummyValue"}}
						fakeCoerce.ToObjectReturns(&objectValue, nil)

						fakeParser := new(FakeParser)
						fakeParser.ParseReturns(identifier, "dummyRefRemainder")

						fakeValueConstructor := new(value.FakeConstructor)
						constructErr := errors.New("constructErr")
						fakeValueConstructor.ConstructReturns(nil, constructErr)

						objectUnderTest := _deReferencer{
							coerce:           fakeCoerce,
							parser:           fakeParser,
							valueConstructor: fakeValueConstructor,
						}

						expectedErr := fmt.Errorf("unable to deReference '%v'; error was %v", providedRef, constructErr.Error())

						/* act */
						_, _, actualErr := objectUnderTest.DeReference(
							providedRef,
							new(model.Value),
						)

						/* assert */
						Expect(actualErr).To(Equal(expectedErr))
					})
				})
				Context("valueConstructor.Construct doesn't err", func() {

					It("should return expected result", func() {

						/* arrange */
						identifier := "identifier"
						refRemainder := "refRemainder"

						fakeCoerce := new(coerce.Fake)
						objectValue := model.Value{Object: map[string]interface{}{identifier: "dummyValue"}}
						fakeCoerce.ToObjectReturns(&objectValue, nil)

						fakeParser := new(FakeParser)
						fakeParser.ParseReturns(identifier, refRemainder)

						fakeValueConstructor := new(value.FakeConstructor)
						identifierValue := model.Value{String: new(string)}
						fakeValueConstructor.ConstructReturns(&identifierValue, nil)

						objectUnderTest := _deReferencer{
							coerce:           fakeCoerce,
							parser:           fakeParser,
							valueConstructor: fakeValueConstructor,
						}

						/* act */
						actualRefRemainder, actualValue, actualErr := objectUnderTest.DeReference(
							"dummyRef",
							new(model.Value),
						)

						/* assert */
						Expect(actualRefRemainder).To(Equal(refRemainder))
						Expect(*actualValue).To(Equal(identifierValue))
						Expect(actualErr).To(BeNil())
					})
				})
			})
		})
	})
})
