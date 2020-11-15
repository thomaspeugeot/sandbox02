import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'


@Component({
  selector: 'app-paragraph-adder',
  templateUrl: './paragraph-adder.component.html',
  styleUrls: ['./paragraph-adder.component.css']
})
export class ParagraphAdderComponent implements OnInit {

	paragraph = {} as ParagraphDB;






  constructor(
    private paragraphService: ParagraphService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.paragraphService.postParagraph( this.paragraph )
    .subscribe(paragraph => {
		this.paragraphService.ParagraphServiceChanged.next("post")
		
		this.paragraph = {} // reset fields
	    console.log("paragraph added")
    });
  }
}
