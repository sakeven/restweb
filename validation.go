package restweb

type ValidationError struct {
	Key, Msg string
}

func (v *ValidationError) String() string {
	return v.Msg
}

type Validation struct {
	HasError  bool
	ValErrors []ValidationError
}

func (v *Validation) Required(obj interface{}, key string) {
	v.Apply(&Required{}, obj, key)
}

func (v *Validation) Min(n int, min int, key string) {
	v.Apply(&Min{min}, n, key)
}

func (v *Validation) Max(n int, max int, key string) {
	v.Apply(&Max{max}, n, key)
}

func (v *Validation) Range(n int, min, max int, key string) {
	v.Apply(&Range{min: min, max: max}, n, key)
}

func (v *Validation) Mail(mail string, key string) {
	pattern := `^\w(\.?\w)*@\w(\.?\w)*\.[A-Za-z]+$`
	v.Apply(&Mail{pattern}, mail, key)
}

func (v *Validation) MinSize(obj interface{}, min int, key string) {
	v.Apply(&MinSize{min}, obj, key)
}

func (v *Validation) MaxSize(obj interface{}, max int, key string) {
	v.Apply(&MaxSize{max}, obj, key)
}

func (v *Validation) Lenth(obj []interface{}, lenth int, key string) {
	v.Apply(&Lenth{lenth}, obj, key)
}

func (v *Validation) Match(obj string, pattern string, key string) {
	v.Apply(&Match{pattern}, obj, key)
}

func (v *Validation) Clear() {
	v.HasError = false
	v.ValErrors = []ValidationError{}
}

func (v *Validation) Apply(vr Validator, obj interface{}, key string) {
	if vr.IsValid(obj) {
		return
	}

	v.HasError = true
	verr := ValidationError{Key: key, Msg: vr.Message()}
	v.ValErrors = append(v.ValErrors, verr)
}

func (v *Validation) RenderErrMap() (ValidError map[string]string) {
	if v.HasError == false {
		return nil
	}
	ValidError = make(map[string]string)
	for _, verr := range v.ValErrors {
		if _, ok := ValidError[verr.Key]; !ok { //render the first key
			ValidError[verr.Key] = verr.String()
		}
	}
	return
}
