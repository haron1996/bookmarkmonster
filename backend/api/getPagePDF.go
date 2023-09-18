package api

/*

	pdfScale := float64(1)
	paperWidth := float64(8.5)
	paperHeight := float64(11)
	marginTop := float64(.4)
	marginBotton := float64(.4)
	marginLeft := float64(.4)
	marginRight := float64(.4)

	sr, err := page.PDF(&proto.PagePrintToPDF{
		Landscape:           false,
		DisplayHeaderFooter: false,
		PrintBackground:     true,
		Scale:               &pdfScale,
		PaperWidth:          &paperWidth,
		PaperHeight:         &paperHeight,
		MarginTop:           &marginTop,
		MarginBottom:        &marginBotton,
		MarginLeft:          &marginLeft,
		MarginRight:         &marginRight,
		PageRanges:          "",
		// HeaderTemplate: ,
		// FooterTemplate: ,
		PreferCSSPageSize: false,
		// TransferMode: ,
	})
	if err != nil {
		panic(err)
	}

	defer sr.Close()

	pdfByte, err := io.ReadAll(sr)
	if err != nil {
		panic(err)
	}

	pdfFile, err := os.Create("page.pdf")
	if err != nil {
		panic(err)
	}

	defer pdfFile.Close()

	_, err = pdfFile.Write(pdfByte)
	if err != nil {
		panic(err)
	}
*/
