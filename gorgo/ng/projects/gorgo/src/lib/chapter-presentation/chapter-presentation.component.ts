import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ChapterDB } from '../chapter-db'
import { ChapterService } from '../chapter.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-chapter-presentation',
	templateUrl: './chapter-presentation.component.html',
	styleUrls: ['./chapter-presentation.component.css']
})
export class ChapterPresentationComponent implements OnInit {

	chapter: ChapterDB;



	ChaptersViaChaptersFieldName = "Chapters"; // Label used to generates the table of Chapter that points to Chapter via Chapters
	ChaptersViaChaptersStructName = "Chapter"; // Label used to generates the table of Chapter that points to Chapter via Chapters
	DocumentsViaChaptersFieldName = "Chapters"; // Label used to generates the table of Document that points to Chapter via Chapters
	DocumentsViaChaptersStructName = "Chapter"; // Label used to generates the table of Document that points to Chapter via Chapters

	constructor(
		private chapterService: ChapterService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getChapter();


		// observable for changes in 
		this.chapterService.ChapterServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getChapter()
					
				}
			}
		)
	}

  getChapter(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.chapterService.getChapter(id)
		.subscribe( 
			chapter => 
			{ 
					this.chapter = chapter
        	}
  	);
  }



	// set presentation outlet
	setPresentationRouterOutlet(structName :string, ID: number) {
		this.router.navigate([{
	  	outlets: {
			presentation: [structName + "-presentation", ID]
	  	}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
	 		outlets: {
	   			editor: ["chapter-detail", ID]
	 	}
   	}]);
 }

}
