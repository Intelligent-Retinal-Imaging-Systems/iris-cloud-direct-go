package Shared

// Model for a person gender
type PersonGender struct {

	// Context of the gender of the following choices
	// BirthGender, IdentityGender
	Context string
	// Value of the Gender of the following choices:
	// U=Unspecified, M=Male, F=Female, N=Non-Binary, O=Other, TM=Transgender male, TF=Transgender female,
	// ND=Non disclosed, EQ=Exploring or Questioning, NL-Not listed, X=Not exclusively Male or Female
	Gender string
}
