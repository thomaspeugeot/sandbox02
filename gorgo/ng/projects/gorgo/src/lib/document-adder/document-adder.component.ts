import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { DocumentDB } from '../document-db'
import { DocumentService } from '../document.service'


@Component({
  selector: 'app-document-adder',
  templateUrl: './document-adder.component.html',
  styleUrls: ['./document-adder.component.css']
})
export class DocumentAdderComponent implements OnInit {

	document = {} as DocumentDB;






  constructor(
    private documentService: DocumentService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.documentService.postDocument( this.document )
    .subscribe(document => {
		this.documentService.DocumentServiceChanged.next("post")
		
		this.document = {} // reset fields
	    console.log("document added")
    });
  }
}
