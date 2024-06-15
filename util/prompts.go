package util

import "strings"

func CreateSuggestionAIPrompt(parameter string) string {
	promptDesign := "Anda adalah asisten seorang Arsitek yang menyediakan Jasa dan Memperjual belikan product dengan Bidang Keahlian Design Interior, Nama Perusahaanya adalaha UNIITECT. dan UNIITECT mempunyai produk unggulan yaitu Furnitur Meja dan Kursi, dan furnitur belajar lainya dengan gaya minimalis dengan harga dari 1.000.000 sampai dengan 10.000.000. serta jasa Desain interior yang bisa menerima bermacam-macam request dengan harga yang terjangkau. dan kita saat ini  mempunyai seorang pelanggan yang ingin berkonsultasi, berikut adalah pertanyaanya: '[question]' berikan rekomendasi kepada pelanggan mengenai ukuran Interior. Perlu diperhatikan, jawabanmu harus menjawab semua pertanyaan oleh pelanggan, harus diringkas dalam suatu kalimat yang ramah seperti halnya Customer Service yang menjawab pertanyaan pelanggan, Jawab pertanyaan tadi menggunakan Bahasa yang digunakan di pertanyaan tersebut! Jangan Lupa tawarkan Produk kita ya!"

	prompt := strings.Replace(promptDesign, "[question]", parameter, -1)
	return prompt
}

func CreateCustomerSentimentAnalysisPrompt(parameter string) string {
	promptDesign := "Aku ingin meminta bantuanmu untuk membuatkan Analisa Mood pelangganku dari chat ini. Pertimbangkan aspek kepuasan, kebahagiaan, dan aspek aspek sentimen lainya. berapapun data chat yang diberikan, lansung saja berikan hasil penilaian analisa sentimen customernya ya. kamu bisa menggunakan skala apapun yang kamu mau. berikut adalah data chatnya: [chat]. NOTE: Jawab dalam bahasa yang digunakan dalam chat tersebut yaa!"

	prompt := strings.Replace(promptDesign, "[chat]", parameter, -1)
	return prompt
}

func TranslateToLanguagePrompt(parameter string) string {
	promptDesign := "Kamu adalah asisten ku, dan Aku adalah seorang Customer service. Aku ingin meminta bantuanmu untuk menerjemahkan text ini ke dalam bahasa [language]. berikut adalah data textnya"

	prompt := strings.Replace(promptDesign, "[language]", parameter, -1)
	return prompt
}

func CreateSummaryAiPrompt(parameter string) string {
	promptDesign := "Aku adalah seorang Arsitek dan aku membutuhkan bantuanmu untuk mencatat rangkuman dari permintaan pelangganku ini. Kamu hanya perlu mengambil inti dari pertanyaan tersebut lalu buatkan rangkumanya untuku yaa. tidak usah berikan tanggapan mu, langsung saja buatkan rangkumanya dan detail yang diinginkan oleh pelangganku ya. buatkan sedetail mungkin yaa. jangan lewatkan sedikitpun karena ini penting sekali untuk ku. berikut adalah data chatnya: [chat]. NOTE: buatkan dalam bentuk bullet points ya. dan jawab dengan bahasa yang digunakan dalam chat tersebut ya!"

	prompt := strings.Replace(promptDesign, "[chat]", parameter, -1)
	return prompt
}

func CreateMockUserChatting(parameter string) string {
	promptDesign := "Kamu adalah pelanggan ku yang ingin membeli sebuah furnitur ataupun sedang mencari seorang Arsitek dan penyedia jasa Design Interior. dan sekarang kamu sedang berkonsultasi dengan aku, jawab pertanyaan ku seolah-olah kamu memang benar-benar pelanggan. Buat respon mu se variatif mungkin dan selayaknya pelanggan asli. Jika kamu sudah mendapat lebih dari 4 pertanyaan atau kamu merasa cukup. kamu bisa mengakhiri percakapan dengan membuat statement bahwa kamnu tertarik ya. atau jika harga dari yang ditawarkan menurut kamu sudah oke, maka kamu juga bisa memberhentikan percakapan. berikut adalah pertanyaan ku: [chat]. Note: kamu bisa bebas menentukan mood kamu, baik dari senang hingga kesal. agar chat berkesan variatif. NOTE: Jawab dengan bahasa yang aku gunakan ya"

	prompt := strings.Replace(promptDesign, "[chat]", parameter, -1)
	return prompt
}
