// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'


// generated table component
@Component({
  selector: 'app-paragraphs-table',
  templateUrl: './paragraphs-table.component.html',
  styleUrls: ['./paragraphs-table.component.css']
})
export class ParagraphsTableComponent implements OnInit {

  // the data source for the table
  paragraphs: ParagraphDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Paragraph
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Content', 'Diagramname', 'Name', 'Structname', 'Edit', 'Delete'];

  constructor(
    private paragraphService: ParagraphService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.paragraphService.ParagraphServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getParagraphs()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getParagraphs()
  }

  getParagraphs(): void {
    if (this.ID == null) {
      this.paragraphService.getParagraphs().subscribe(
        Paragraphs => {
          this.paragraphs = Paragraphs;
        }
      )
    }
  
  }

  // newParagraph initiate a new paragraph
  // create a new Paragraph objet
  newParagraph() {
  }

  deleteParagraph(paragraphID: number, paragraph: ParagraphDB) {
    // la liste des paragraphs est amputée du paragraph avant le delete afin
    // de mettre à jour l'IHM
    this.paragraphs = this.paragraphs.filter(h => h !== paragraph);

    this.paragraphService.deleteParagraph(paragraphID).subscribe();
  }

  editParagraph(paragraphID: number, paragraph: ParagraphDB) {

  }

  // display paragraph in router
  displayParagraphInRouter(paragraphID: number) {
    this.router.navigate( ["paragraph-display", paragraphID])
  }

  // set editor outlet
  setEditorRouterOutlet(paragraphID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["paragraph-detail", paragraphID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(paragraphID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["paragraph-presentation", paragraphID]
      }
    }]);
  }
}
