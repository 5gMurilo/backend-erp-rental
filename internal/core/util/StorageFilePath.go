package util

func NewPathFile(employeeName string, fileType string, fileName string) string {
	var url string
	//if fileType == "outros" {
	//	url = "https://graph.microsoft.com/v1.0/me/drive/items/root:/Rede - RH/RH - América Rental/sistema/" + employeeName + "/" + fileName + ":/content"
	//} else {
	//	url = "https://graph.microsoft.com/v1.0/me/drive/items/root:/Rede - RH/RH - América Rental/sistema/" + employeeName + "/" + fileType + "/" + fileName + ":/content"
	//}

	if fileType == "outros" {
		url = "https://graph.microsoft.com/v1.0/me/drive/items/root:/Attachments/" + employeeName + "/" + fileName + ":/content"
	} else {
		url = "https://graph.microsoft.com/v1.0/me/drive/items/root:/Attachments/" + employeeName + "/" + fileType + "/" + fileName + ":/content"
	}
	return url
}
