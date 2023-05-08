package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	DataSplit := make([][]string, len(data))
	for i := 0; i < len(data); i++ {
		DataSplit[i] = append(DataSplit[i], strings.Split(data[i], "-")...)
	}
	KeyDict := []string{}
	rev := ""
	for _, key := range DataSplit[:] {
		if rev != key[0] {
			rev = key[0]
			KeyDict = append(KeyDict, key[0])
		}
	}
	// fmt.Println(DataSplit)
	Hasil:=make(map[string][]string, len(KeyDict))
	for key:=0; key<len(KeyDict); key++{
		for i:=0; i<len(DataSplit); i++{
			for j:=i+1; j<len(DataSplit); j++{
				if DataSplit[i][0]==KeyDict[key] && DataSplit[j][0]==KeyDict[key]{
					if j%2==0{
						break
					}
					if DataSplit[i][0]=="phone" &&  DataSplit[j][0]=="phone"{
						Hasil[KeyDict[key]]=append(Hasil[KeyDict[key]], DataSplit[i][3])
						Hasil[KeyDict[key]]=append(Hasil[KeyDict[key]], DataSplit[j][3])
						break
					}
					itemString:=DataSplit[i][3]+" "+DataSplit[j][3]
					n:=0
					for _, SameName:=range Hasil[KeyDict[key]]{
						Check:=false
						Check=strings.Contains(SameName, itemString)
						if Check{
							n++
							break
						}
					}
					if n==1{break}
					Hasil[KeyDict[key]]=append(Hasil[KeyDict[key]], itemString)
					break
				}
			}
		}
	}
	return Hasil // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"data1-0-first-hADsbQmeBaThwno", "data1-0-last-opTgtLxiWsTlwoU", "data2-0-first-UWxzjUqaaAIxyXN", "data2-0-last-FBFnxFuMAXdfpyH", "data3-0-first-MPcdKXSFoesHMvj", "data3-0-last-HHLphaxkihBCMce", "data4-0-first-jLfUTYtOWmbpmGi", "data4-0-last-fqaOBphLGVSbuiA", "data5-0-first-skzZzTISktBlRfz", "data5-0-last-DGWvrMescAGqgld", "data6-0-first-vChqqqCWFoNCENK", "data6-0-last-pesZrfxICioATvN", "data7-0-first-EPqcJJRariXAjMv", "data7-0-last-GAdCVutAfDwtxou", "data8-0-first-iPxtWrcaICLLhiD", "data8-0-last-rbpEIDhtdrojPdT", "data1-1-first-OzPymccCQfgMqmc", "data1-1-last-ZZElFjmTACZQAOv", "data2-1-first-WNggUUhRmpDohtf", "data2-1-last-uJGrYKdEGOFpRRC", "data3-1-first-IyfBXsRSGXuhXfo", "data3-1-last-lAOoVGRhnitQfAn", "data4-1-first-BudbOglTAoxkSyh", "data4-1-last-YibpJpwBENHkFlk", "data5-1-first-fSagkbEPAMdmhpZ", "data5-1-last-PsLDTFFuFDkTnuh", "data6-1-first-TKoiDLdjXqWmZUh", "data6-1-last-eMvMKITanhcPJAn", "data7-1-first-QvLjiZPnjXMeJqX", "data7-1-last-LPqxhLSRJcbBXgL", "data8-1-first-snaBbwNYGlwmOpb", "data8-1-last-vBpLrZIXlEjFqIv", "data1-2-first-ePsSyzjsBpyZBsX", "data1-2-last-ashcCpaeEuzErER", "data2-2-first-PgVzbitcASicfTF", "data2-2-last-SUcGNvhrkTZbemD", "data3-2-first-qLzIqhAIGNScNzr", "data3-2-last-pRYUtGsuuKZiYjS", "data4-2-first-NUuSPFmmCLknzEw", "data4-2-last-RXOJwHVTimJzuUT", "data5-2-first-YBHTCsHRUjIqhGf", "data5-2-last-aXpKixfHAITNcnx", "data6-2-first-pXRNRzqZlExSUvQ", "data6-2-last-sTkoVsWTbiTmTsf", "data7-2-first-hawuMHQWDOiCxcZ", "data7-2-last-kzYBsovILRwtBAJ", "data8-2-first-nOegzOmivhMMtXT", "data8-2-last-bIkXXrhstgupPnb", "data1-3-first-QMTyOFgdprsLqzZ", "data1-3-last-CiwcMnWFRoLnJan", "data2-3-first-cHCnEBmzbRxXlEE", "data2-3-last-LRBmoyOZgoCAHAV", "data3-3-first-niVgsQHHIlaLKOt", "data3-3-last-lBDjPpHswrhSWbt", "data4-3-first-LtlyjTbOScmPVzT", "data4-3-last-pqHCHcAhVGoDNYO", "data5-3-first-tzcXwhMLCHbuoGF", "data5-3-last-JXkInnzLFlnnEKF", "data6-3-first-MHrIUPIcbueFgwV", "data6-3-last-SYlMYYaplHZKAwA", "data7-3-first-WiNxQYlVmJWFoEB", "data7-3-last-ySphKoQpvgAOLeo", "data8-3-first-CWyZuBCdSzZCBWR", "data8-3-last-RiFheGvmzvwzCuy", "data1-4-first-ADPkmzdDBqOHyVt", "data1-4-last-NkcRgLKOHvgaXXN", "data2-4-first-EoHZNUuXvWKcnwF", "data2-4-last-nDMClcpcSvyhAwN", "data3-4-first-VCnhxRlYASFYXUk", "data3-4-last-MWFGMmEefWopoqA", "data4-4-first-nXmJluArdcrBKme", "data4-4-last-zroGsjTYBLLQKvO", "data5-4-first-VDDqjMnkYeepmPl", "data5-4-last-PrgSgaKBizgtKYN", "data6-4-first-vrUrmNRyESVcWFB", "data6-4-last-XsaGGxBGtlXtXYZ", "data7-4-first-XcNOirUgjkOrKpN", "data7-4-last-crZbmwActlzYjUl", "data8-4-first-jMozONrVUkapaLV", "data8-4-last-jcKreygKUoNTZrZ", "data1-5-first-VLVjqmufPeUjubI", "data1-5-last-HjyXCqPcFZQMBpE", "data2-5-first-jWZrhHSJfcLNWGA", "data2-5-last-GSjryISdTkZXtVL", "data3-5-first-TKwLKRhvbcKEcaH", "data3-5-last-eCWXqqpclUVlRZO", "data4-5-first-GCBTtTwtTYbRoED", "data4-5-last-FbVMRUhuKNCOdoi", "data5-5-first-fjksmhqOidSOMam", "data5-5-last-ajjwVSUfZIZBbKm", "data6-5-first-GPDwtpAxzdMmsQp", "data6-5-last-UleZpaTZxvcvnjd", "data7-5-first-aFRYhTESsUnvYAZ", "data7-5-last-BZjkRsePMJVinBk", "data8-5-first-tsfFrcoDeWWkkEi", "data8-5-last-tXQBvxjglmGZNhR", "data1-6-first-kXedHNMQDekRszJ", "data1-6-last-ttlhfqaqwQbmvHR", "data2-6-first-YMWNbUfTizcxMfx", "data2-6-last-VsNxvFcUomGdfmU", "data3-6-first-TLHOXCawfrhKPJQ", "data3-6-last-tkwNfXHFPylBPPX", "data4-6-first-gYGqTveCRqARygb", "data4-6-last-jkBhXjWsdIbCLKz", "data5-6-first-lgGtfsVhfLthaNf", "data5-6-last-McYdLNdmluqxOcz", "data6-6-first-UkPuTtcqpeAzioT", "data6-6-last-fnRKTSlssSjqdvS", "data7-6-first-lieKCgPYKycLama", "data7-6-last-vnoCitfCEGLvhHm", "data8-6-first-PrbOjOSePXxBArq", "data8-6-last-CzECSXifgLEvRrV", "data1-7-first-QXKeiXVUOWsropZ", "data1-7-last-vHNCubmgUPsYNhg", "data2-7-first-uHmSKophvqWCUew", "data2-7-last-ZeISaNaTramfUfU", "data3-7-first-DZnKGFiHBPCgJTn", "data3-7-last-btNfWywhEsYCTIn", "data4-7-first-ysMHEzLdWoMAsTg", "data4-7-last-HChHCxkyBDjuzVW", "data5-7-first-pURnYFyyFXmqcWv", "data5-7-last-vnPYYucXLJqvktv", "data6-7-first-JUMfsWZVxQhuPGv", "data6-7-last-iKuzEQrFcAoCvML", "data7-7-first-bmNgIBFpJgPLUYb", "data7-7-last-KemgoZLrfIQPnAu", "data8-7-first-BCwQwsYAhnjwXnC", "data8-7-last-tBgIqeOMwjRZJKu", "data1-8-first-qeEunZQyIuBJnMw", "data1-8-last-ylkeVFfeokGinXB", "data2-8-first-NLEaqyUvvVRSsUk", "data2-8-last-xTlresEaNEAhoPa", "data3-8-first-JWTzCcdbTAhZXdd", "data3-8-last-tdrHcQpwsbcmitW", "data4-8-first-UxGPwNhEqMPVxMp", "data4-8-last-rVqzIesiBumDEeK", "data5-8-first-RcZXwHwWjHJgtBF", "data5-8-last-AsfUcuvEbswgoKp", "data6-8-first-wHEwsHgoWJVnxwt", "data6-8-last-NgbvowIpIDIYjUg", "data7-8-first-eKPccVtdqZUoJBV", "data7-8-last-wqmwMJxmnEbCliB", "data8-8-first-EyDYEmoPtuRkseV", "data8-8-last-floINzEgMfUWvsU", "data1-9-first-CQsBmEkNCJYLcIo", "data1-9-last-UUcgyVIqKmsXvIs", "data2-9-first-XAigOyKeVMCAkdE", "data2-9-last-DsxptLLDEFfjPFA", "data3-9-first-YFQIpLKYBWajvJf", "data3-9-last-TklxyZuZZqsIDLr", "data4-9-first-TqBRxwlJNqUGMfT", "data4-9-last-iVnJOSPWPzHrUnH", "data5-9-first-eROZuqYfRQaGMVT", "data5-9-last-FtfCxLGoIppfLFH", "data6-9-first-jVGyKXPIoPWYsmr", "data6-9-last-aBrBqbKmXowVboD", "data7-9-first-VZVyiQjqETPavOa", "data7-9-last-fKTtdahwjDXHopm", "data8-9-first-iuOCSmmBHEOVOAp", "data8-9-last-SjxfDuUkSOZGvwq", "data1-10-first-RYDucDHRQFVdouW", "data1-10-last-dnRXHgPmfFVyPma", "data2-10-first-GEbGzFanLtfKXCl", "data2-10-last-NYhaXPCUxWPIeBG", "data3-10-first-BwirXTlxcqeWCop", "data3-10-last-sExDHMamCdQafJR", "data4-10-first-ekAnzhEXxazoeXo", "data4-10-last-zkeHBpinxkIKSRx", "data5-10-first-JFlchOxQXRXMoMt", "data5-10-last-ZBrnxCqtSfOVLyt", "data6-10-first-PhBHxEbexsSpzfY", "data6-10-last-lVGsBvyrdHgTyrM", "data7-10-first-wJVopESwHvbNOto", "data7-10-last-HJMAlVJFRMAJLVg", "data8-10-first-jxNAdCJVMFLyVRk", "data8-10-last-ojWTpDsXUbZJTHX", "data1-11-first-ZJvwdiSWTdKHnTc", "data1-11-last-npRQFUlfpzRcqyc", "data2-11-first-DNicVkhMTHgQTtP", "data2-11-last-TGgtuSXsfhkRqmA", "data3-11-first-cewNqCkemGLrxgn", "data3-11-last-hKHRvaDGhaGOLjY", "data4-11-first-fuoxnbdGrEGHmEz", "data4-11-last-LOQSNbYezxdnLjl", "data5-11-first-buqFdQCfqYQYUpY", "data5-11-last-XwDtFBGTfdETIOl", "data6-11-first-fMueBluAQDLMjpc", "data6-11-last-ysFizGPAmNcQKux", "data7-11-first-VRyCPbXdJEPiVJV", "data7-11-last-RIRnnKiGfsNvMPx", "data8-11-first-ewHCjEKTopRmxaE", "data8-11-last-KhEKPwZUyRsTAvE", "data1-12-first-BVyfvZhscwHYJos", "data1-12-last-exxHvSoqKJUGsMo", "data2-12-first-YLxPDGzBfvKNQcF", "data2-12-last-KtSLLXzquwyYFug", "data3-12-first-VxMXIMCrgSpalWQ", "data3-12-last-ExHXUIbYuPpLEIc", "data4-12-first-kJsztKdVRpzZcxk", "data4-12-last-JRCPBripVOSGPIo", "data5-12-first-RuxkpAczTnirsSL", "data5-12-last-gxjUtNgzqsxcUFm", "data6-12-first-dxsHJNjkZyMhfBS", "data6-12-last-lQvrAxXCEnBmovO", "data7-12-first-smLUUXMJbwfhEir", "data7-12-last-YOdQUrdOBxqGUzj", "data8-12-first-vWKosXurkmWOFCe", "data8-12-last-hWFVdvkQAgFxQlq", "data1-13-first-UjbAsJVNfxUeMTp", "data1-13-last-sOsJcRLqRvlivQe", "data2-13-first-jpGdwuEjYMfpyqV", "data2-13-last-fVutTyInEKKWBTq", "data3-13-first-OcphTjZfqvHUfXS", "data3-13-last-QduGCsHosZycjAt", "data4-13-first-LtQdeFoBfSsugkp", "data4-13-last-aBjpCCCOnRMAbpd", "data5-13-first-XZkhWPznenUUIuA", "data5-13-last-bpiimIMEQCqFCWU", "data6-13-first-nCSSIftkOwzfUCp", "data6-13-last-KZdOcZdRPzlMoQN", "data7-13-first-ZNeqgmQKpRYnjAt", "data7-13-last-fyQqMHojoADPWhA", "data8-13-first-fKHSIIPguwUNYHg", "data8-13-last-yuJGAjAJBMejTmL", "data1-14-first-tHfvMYaUVguKQft", "data1-14-last-xqSbMFBqKweeQVY", "data2-14-first-sigvYjCLeWwuHzw", "data2-14-last-mOXbgyAXEOPcMmL", "data3-14-first-iaIHkHsYENACzBY", "data3-14-last-FuWzokbZWKajslH", "data4-14-first-GSNPIYYytNhzPNT", "data4-14-last-QZRInMoOYxykfbH", "data5-14-first-GvJYqJkqPHMicvL", "data5-14-last-UjNHubDUOQTtXrT", "data6-14-first-AjLEOTDNxzYSGZH", "data6-14-last-HHjUFUeGTJKeHyD", "data7-14-first-hUOwJizbjAplidH", "data7-14-last-JjChoZFPtmycXyj", "data8-14-first-IXMAYWkzwlIPoKr", "data8-14-last-DfekEDRsHxXMrfC", "data1-15-first-gbRpgXZkrPTeowg", "data1-15-last-TQZwcbytBsCrMjb", "data2-15-first-NXGlIsMWDXhzjgj", "data2-15-last-nEDlQWZDWDWipKT", "data3-15-first-XkidonWguwomDCr", "data3-15-last-ziZJuCqMYdtFAmv", "data4-15-first-TgEeenLJPuzuxEz", "data4-15-last-kwWBWJutiFuikgC", "data5-15-first-LgAheemntptBoMm", "data5-15-last-WcqiJuQEECKRFwS", "data6-15-first-OkniGQxACGkObOq", "data6-15-last-fBmOSIAyFJDJYfa", "data7-15-first-YWFDTgShdDcxFyf", "data7-15-last-zNzvSLQDLQNGMOh", "data8-15-first-WdXfyYLkfwXnTWR", "data8-15-last-eJaIqNXoeluBAqB", "data1-16-first-qbcUbmkcOIpbscp", "data1-16-last-BnuViTQvcRJQzur", "data2-16-first-eDCRWIzlKsoZVvr", "data2-16-last-mhvZPEyoOyXJzVD", "data3-16-first-lmEtTNjBoBcyLor", "data3-16-last-iUxqbInFPsSFwYE", "data4-16-first-NzgdDBUmrfJxLko", "data4-16-last-FtLiuxMRllGrBNY", "data5-16-first-xuYmiqZCYwgikmk", "data5-16-last-IMoeqTTSgjVFlOV", "data6-16-first-McLGOzqHybrgeUd", "data6-16-last-kETPsIUTsQIsSZV", "data7-16-first-jxxCAxHIEJnXhsy", "data7-16-last-QGwvCtdojrPqcgS", "data8-16-first-RLmAiyzpqeioHQi", "data8-16-last-AnZfZjNBAGZRCCt", "data1-17-first-VUlACXbRURdcEiW", "data1-17-last-HXJwyIGDrMVEfOa", "data2-17-first-cxdosXQSRmEqYeH", "data2-17-last-jhAoEkvfgnDuJSm", "data3-17-first-oVUFcQZvbRuxlwf", "data3-17-last-PzMhTdsziBvxVQQ", "data4-17-first-vvzMrvzzcQuPCUU", "data4-17-last-jMzlSrJTQXjUIsx", "data5-17-first-XAIVOjokKcOlGzM", "data5-17-last-GSrmNryQmWMslgp", "data6-17-first-OCelBQrtWeUAUKp", "data6-17-last-qDvStHDFYPbxUOi", "data7-17-first-jHRybwsOivguVzh", "data7-17-last-RwrPmokWXJSByxm", "data8-17-first-DomGaxZiBuOSMgm", "data8-17-last-nSSSYqjujrqSisz", "data1-18-first-PqNDpWoSnRtlwbv", "data1-18-last-OGEvlcYjxVZhukz", "data2-18-first-XjgDDzkLUiIebjw", "data2-18-last-SloUDbEiIVVltUz", "data3-18-first-mZlfXEIbHwbGtrb", "data3-18-last-WbBIiyJEgeYnZrc", "data4-18-first-GwZIVlfjRkzJoPt", "data4-18-last-AJVQnQOnnCMEwPi", "data5-18-first-FlPEMWJmNpyShGa", "data5-18-last-gNKLwaraUYipoWL", "data6-18-first-YWOycDRfHaprHtQ", "data6-18-last-QtDNwJiviSIcAcz", "data7-18-first-YmwgudjwkEsPAIk", "data7-18-last-odYrmBsygvPHhMR", "data8-18-first-NYduHnp"}
	res := ChangeOutput(data)
	for key, data:=range res{
		fmt.Println(key , ":", data)
	}
	fmt.Println(res)
}
