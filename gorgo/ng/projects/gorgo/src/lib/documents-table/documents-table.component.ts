// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { DocumentDB } from '../document-db'
import { DocumentService } from '../document.service'


// generated table component
@Component({
  selector: 'app-documents-table',
  templateUrl: './documents-table.component.html',
  styleUrls: ['./documents-table.component.css']
})
export class DocumentsTableComponent implements OnInit {

  // the data source for the table
  documents: DocumentDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Document
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Author1', 'Author1function', 'Author2', 'Author2function', 'Authority', 'Authorityfunction', 'Date', 'Name', 'Outputfilepath', 'Templatefilepath', 'Title', 'Verifier', 'Verifierfunction', 'Edit', 'Delete'];

  constructor(
    private documentService: DocumentService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.documentService.DocumentServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getDocuments()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getDocuments()
  }

  getDocuments(): void {
    if (this.ID == null) {
      this.documentService.getDocuments().subscribe(
        Documents => {
          this.documents = Documents;
        }
      )
    }
  
  }

  // newDocument initiate a new document
  // create a new Document objet
  newDocument() {
  }

  deleteDocument(documentID: number, document: DocumentDB) {
    // la liste des documents est amputée du document avant le delete afin
    // de mettre à jour l'IHM
    this.documents = this.documents.filter(h => h !== document);

    this.documentService.deleteDocument(documentID).subscribe();
  }

  editDocument(documentID: number, document: DocumentDB) {

  }

  // display document in router
  displayDocumentInRouter(documentID: number) {
    this.router.navigate( ["document-display", documentID])
  }

  // set editor outlet
  setEditorRouterOutlet(documentID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["document-detail", documentID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(documentID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["document-presentation", documentID]
      }
    }]);
  }
}
