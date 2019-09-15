package model

import (
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"go-admin/verification-code/code-server/store"
	"time"
)

var defaultExpirytime = 120000 * time.Millisecond
var configD = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 6,
}

//config struct for audio
//声音验证码配置
var configA = base64Captcha.ConfigAudio{
	CaptchaLen: 6,
	Language:   "zh",
}

//config struct for Character
//字符,公式,验证码配置
var configC = base64Captcha.ConfigCharacter{
	Height: 60,
	Width:  240,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeArithmetic,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}

func Init(cli *redis.Client, prefix string, exxpirytime time.Duration) {
	if exxpirytime == 0 {
		exxpirytime = defaultExpirytime
	}
	base64Captcha.SetCustomStore(store.NewRedisStore(cli, prefix, exxpirytime))
}
func buildConfigDigit(w ,h int64) base64Captcha.ConfigDigit {
	return base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 6,
	}
}
func buildConfigConfigCharacter(w,h int64) base64Captcha.ConfigCharacter  {
	return base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
}
func CreateCode(t int64,w int64,h int64) (string, string) {
	switch t {
	case 1:
		config := buildConfigDigit(w,h)
		idKey, capD := base64Captcha.GenerateCaptcha("", config)
		//以base64编码
		base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
		return idKey, base64stringD
	case 2:
		config := buildConfigConfigCharacter(w,h)
		idKeyC, capC := base64Captcha.GenerateCaptcha("", config)
		//write to base64 string.
		base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
		return idKeyC, base64stringC
	case 3:
		idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
		//write to base64 string.
		//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
		base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
		return idKeyA, base64stringA
	default:
		return "", ""
	}
}
func CheckCode(key, verifyValue string) bool {
	return base64Captcha.VerifyCaptchaAndIsClear(key, verifyValue, true)
}
