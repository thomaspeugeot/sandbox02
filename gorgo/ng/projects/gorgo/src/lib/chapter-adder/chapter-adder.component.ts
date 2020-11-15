import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ChapterDB } from '../chapter-db'
import { ChapterService } from '../chapter.service'


@Component({
  selector: 'app-chapter-adder',
  templateUrl: './chapter-adder.component.html',
  styleUrls: ['./chapter-adder.component.css']
})
export class ChapterAdderComponent implements OnInit {

	chapter = {} as ChapterDB;






  constructor(
    private chapterService: ChapterService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.chapterService.postChapter( this.chapter )
    .subscribe(chapter => {
		this.chapterService.ChapterServiceChanged.next("post")
		
		this.chapter = {} // reset fields
	    console.log("chapter added")
    });
  }
}
