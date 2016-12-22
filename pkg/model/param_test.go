package model

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/util/format"
)

var _ = Describe("Param", func() {

  Context("when formatting to/from yaml", func() {
    yaml := format.NewYamlFormat()

    Context("with non-nil $.dir", func() {

      It("should have expected attributes", func() {

        /* arrange */
        expectedParam := Param{
          Dir:&DirParam{
            Name:"dummyName",
            Description:"dummyDescription",
            IsSecret:true,
          },
        }

        /* act */
        providedYaml, err := yaml.From(expectedParam)
        if (nil != err) {
          panic(err)
        }

        actualParam := Param{}
        yaml.To(providedYaml, &actualParam)

        /* assert */
        Expect(actualParam).To(Equal(expectedParam))

      })

    })

    Context("with non-nil $.file", func() {

      It("should have expected attributes", func() {

        /* arrange */
        expectedParam := Param{
          File:&FileParam{
            Name:"dummyName",
            Description:"dummyDescription",
            IsSecret:true,
          },
        }

        /* act */
        providedYaml, err := yaml.From(expectedParam)
        if (nil != err) {
          panic(err)
        }

        actualParam := Param{}
        yaml.To(providedYaml, &actualParam)

        /* assert */
        Expect(actualParam).To(Equal(expectedParam))

      })

    })

    Context("with non-nil $.netSocket", func() {

      It("should have expected attributes", func() {

        /* arrange */
        expectedParam := Param{
          NetSocket:&NetSocketParam{
            Name:"dummyName",
            Description:"dummyDescription",
            IsSecret:true,
          },
        }

        /* act */
        providedYaml, err := yaml.From(expectedParam)
        if (nil != err) {
          panic(err)
        }

        actualParam := Param{}
        yaml.To(providedYaml, &actualParam)

        /* assert */
        Expect(actualParam).To(Equal(expectedParam))

      })

    })

    Context("with non-nil $.string", func() {

      It("should have expected attributes", func() {

        /* arrange */
        expectedParam := Param{
          String: &StringParam{
            Default:"dummyDefault",
            Description:"dummyDescription",
            Constraints:&StringConstraints{
              Length: &StringLengthConstraint{
                Min: 0,
                Max: 1000,
                Description:"dummyStringLengthConstraintDescription",
              },
              Patterns:[]*StringPatternConstraint{
                {
                  Regex:".*",
                  Description:"dummyStringPatternConstraintDescription",
                },
              },
            },
            Name:"dummyName",
            IsSecret:true,
          },
        }

        /* act */
        providedYaml, err := yaml.From(expectedParam)
        if (nil != err) {
          panic(err)
        }

        actualParam := Param{}
        yaml.To(providedYaml, &actualParam)

        /* assert */
        Expect(actualParam).To(Equal(expectedParam))

      })

    })

  })

})