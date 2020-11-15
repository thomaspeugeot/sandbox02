package walk

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgSplitterTemplateTS ...
const NgSplitterTemplateTS = `import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-{{pkgname}}-splitter',
  templateUrl: './splitter.component.html',
  styleUrls: ['./splitter.component.css']
})
export class SplitterComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
}
`

// NgSplitterTemplateHTML ...
const NgSplitterTemplateHTML = `<div style="width: 100%; height: 100%; background: grey(16);">
<as-split direction="horizontal">
	<as-split-area size="20" maxSize="30">
		<as-split direction="vertical">
			<as-split-area>
				<app-sidebar></app-sidebar>
			</as-split-area>
		</as-split>
	</as-split-area>
	<as-split-area size="60" maxSize="100">
		<as-split direction="vertical">
		<as-split-area>
			<router-outlet name="table"></router-outlet>
		</as-split-area>
		<as-split-area>
			<router-outlet name="presentation"></router-outlet>
		</as-split-area>
</as-split>
	</as-split-area>
	<as-split-area size="20" maxSize="100">
		<as-split direction="vertical">
			<as-split-area>
				<router-outlet name="editor"></router-outlet>
			</as-split-area>
		</as-split>
	</as-split-area>
</as-split>
</div>
`

// NgSplitterTemplateCSS ...
const NgSplitterTemplateCSS = ``

// GenNgSplitter generates the splitter on the front
func GenNgSplitter(db *gorm.DB) {

	fileTS, fileHTML, fileCSS := createDirAndTreeFilesInNgTargetPath("splitter", "")
	defer fileTS.Close()
	defer fileHTML.Close()
	defer fileCSS.Close()

	res := NgSplitterTemplateTS
	res = strings.ReplaceAll(res, "{{pkgname}}", strings.ToLower(PkgName))
	fmt.Fprint(fileTS, res)

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	LISTOFSTRUCT := ""

	for _, _struct := range structs {
		LISTOFSTRUCT += strings.ReplaceAll(listItemTemplate, "{{Structname}}", _struct.Name)
	}
	fmt.Fprint(fileHTML, strings.ReplaceAll(NgSplitterTemplateHTML, "{{LISTOFSTRUCT}}", LISTOFSTRUCT))

	fmt.Fprint(fileCSS, NgSplitterTemplateCSS)
}
