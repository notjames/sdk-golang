package interpolate

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/pkg/model"
)

var _ = Describe("Interpolate", func() {
	Context("New()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(New()).Should(Not(BeNil()))
		})
	})
	Context("passed scope containing string", func() {
		It("should call stringInterpolater.Interpolate w/ expected args & return result", func() {
			/* arrange */
			providedTemplate := "dummyTemplate"
			providedVarName := "dummyVarName"
			providedVarData := "dummyVarData"
			providedScope := map[string]*model.Data{
				providedVarName: {String: providedVarData},
			}

			expectedTemplateArg := providedTemplate
			expectedVarName := providedVarName
			expectedVarData := providedVarData
			expectedResult := "dummyResult"

			fakeStringInterpolater := new(fakeStringInterpolater)
			fakeStringInterpolater.InterpolateReturns(expectedResult)

			objectUnderTest := interpolate{
				stringInterpolater: fakeStringInterpolater,
			}

			/* act */
			actualResult := objectUnderTest.Interpolate(providedTemplate, providedScope)

			/* assert */
			actualTemplateArg, actualVarName, actualVarData :=
				fakeStringInterpolater.InterpolateArgsForCall(0)

			Expect(actualTemplateArg).To(Equal(expectedTemplateArg))
			Expect(actualVarName).To(Equal(expectedVarName))
			Expect(actualVarData).To(Equal(expectedVarData))
			Expect(actualResult).To(Equal(expectedResult))
		})
	})

	Context("passed scope containing number", func() {
		It("should call numberInterpolater.Interpolate w/ expected args & return result", func() {
			/* arrange */
			providedTemplate := "dummyTemplate"
			providedVarName := "dummyVarName"
			providedVarData := 1.2
			providedScope := map[string]*model.Data{
				providedVarName: {Number: providedVarData},
			}

			expectedTemplateArg := providedTemplate
			expectedVarName := providedVarName
			expectedVarData := providedVarData
			expectedResult := "dummyResult"

			fakeNumberInterpolater := new(fakeNumberInterpolater)
			fakeNumberInterpolater.InterpolateReturns(expectedResult)

			objectUnderTest := interpolate{
				numberInterpolater: fakeNumberInterpolater,
			}

			/* act */
			actualResult := objectUnderTest.Interpolate(providedTemplate, providedScope)

			/* assert */
			actualTemplateArg, actualVarName, actualVarData :=
				fakeNumberInterpolater.InterpolateArgsForCall(0)

			Expect(actualTemplateArg).To(Equal(expectedTemplateArg))
			Expect(actualVarName).To(Equal(expectedVarName))
			Expect(actualVarData).To(Equal(expectedVarData))
			Expect(actualResult).To(Equal(expectedResult))
		})
	})
})
