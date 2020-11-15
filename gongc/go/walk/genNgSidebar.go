package walk

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgSidebarTemplateTS ...
const NgSidebarTemplateTS = `import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  constructor(
	private router: Router,
  ) { }

  ngOnInit(): void {
  }

  setTableRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        table: [path]
      }
    }]);
  }
  
  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        editor: [path]
      }
    }]);
  }
}
`

// NgSidebarTemplateHTML ...
const NgSidebarTemplateHTML = `<mat-list>
{{LISTOFSTRUCT}}</mat-list>
`

const listItemTemplate = `	<mat-list-item 
		class="row-link">
		<span (click)="setTableRouterOutlet( '{{structname}}s' )" >{{Structname}}s &nbsp;</span>
		<mat-icon (click)="setEditorRouterOutlet( '{{structname}}-adder' )">
			add_circle_outline
		</mat-icon>
		</mat-list-item>
`

// NgSidebarTemplateCSS ...
const NgSidebarTemplateCSS = `.mat-row-link{
	position: absolute;
	width: 100%;
	height: 100%;
	left: 0;
	top: 0;          
	   
  }
  
  .row-link:hover {
	background-color: rgba(0, 0, 0, .05);
	cursor: pointer;
  }

  mat-icon:hover{
    background-color: rgba(0, 0, 0, .05);
    cursor: pointer;
  }
`

// GenNgSidebar generates the sidebar on the front
func GenNgSidebar(db *gorm.DB) {

	fileTS, fileHTML, fileCSS := createDirAndTreeFilesInNgTargetPath("sidebar", "")
	defer fileTS.Close()
	defer fileHTML.Close()
	defer fileCSS.Close()

	fmt.Fprint(fileTS, NgSidebarTemplateTS)

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	LISTOFSTRUCT := ""

	for _, _struct := range structs {

		res := strings.ReplaceAll(listItemTemplate, "{{Structname}}", _struct.Name)
		res = strings.ReplaceAll(res, "{{structname}}", strings.ToLower(_struct.Name))
		LISTOFSTRUCT += res
	}
	fmt.Fprint(fileHTML, strings.ReplaceAll(NgSidebarTemplateHTML, "{{LISTOFSTRUCT}}", LISTOFSTRUCT))

	fmt.Fprint(fileCSS, NgSidebarTemplateCSS)
}
