import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-paragraph-presentation',
	templateUrl: './paragraph-presentation.component.html',
	styleUrls: ['./paragraph-presentation.component.css']
})
export class ParagraphPresentationComponent implements OnInit {

	paragraph: ParagraphDB;



	ChaptersViaParagraphsFieldName = "Paragraphs"; // Label used to generates the table of Chapter that points to Paragraph via Paragraphs
	ChaptersViaParagraphsStructName = "Paragraph"; // Label used to generates the table of Chapter that points to Paragraph via Paragraphs

	constructor(
		private paragraphService: ParagraphService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getParagraph();


		// observable for changes in 
		this.paragraphService.ParagraphServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getParagraph()
					
				}
			}
		)
	}

  getParagraph(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.paragraphService.getParagraph(id)
		.subscribe( 
			paragraph => 
			{ 
					this.paragraph = paragraph
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
	   			editor: ["paragraph-detail", ID]
	 	}
   	}]);
 }

}
