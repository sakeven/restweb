package restweb

type ValidationError struct {
	msg string
}

func (v *ValidationError) String() string {
	return v.msg
}

type Validation struct {
	HasError  bool
	ValErrors []ValidationError
}

func (v *Validation) Required(obj interface{}) {
	v.Apply(Required{}, obj)
}

func (v *Validation) Min(n int, min int) {
	v.Apply(Min{min}, n)
}

func (v *Validation) Max(n int, max int) {
	v.Apply(Max{max}, n)
}

func (v *Validation) Range(n int, min, max int) {
	v.Apply(Range{min: min, max: max}, n)
}
func (v *Validation) Mail(mail string) {
	pattern := `^\w(\.?\w)*@\w(\.?\w)*\.[A-Za-z]+$`
	v.Apply(Match{pattern}, mail)
}

// func (v *Validation) Lenth(obj []interface{}, length int) {
// 	if len(obj) != length {
// 		v.HasError = true
// 	}
// }

func (v *Validation) Match(obj string, pattern string) {
	v.Apply(Match{pattern}, obj)
}

func (v *Validation) Clear() {
	v.HasError = false
}

func (v *Validation) Format(format string) {

}

func (v *Validation) Apply(vr Validator, obj interface{}) {
	if vr.IsValid(obj) {
		return
	}

	v.HasError = true
}
