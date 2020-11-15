import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-field-presentation',
	templateUrl: './field-presentation.component.html',
	styleUrls: ['./field-presentation.component.css']
})
export class FieldPresentationComponent implements OnInit {

	field: FieldDB;



	ClassshapesViaFieldsFieldName = "Fields"; // Label used to generates the table of Classshape that points to Field via Fields
	ClassshapesViaFieldsStructName = "Field"; // Label used to generates the table of Classshape that points to Field via Fields

	constructor(
		private fieldService: FieldService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getField();


		// observable for changes in 
		this.fieldService.FieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getField()
					
				}
			}
		)
	}

  getField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.fieldService.getField(id)
		.subscribe( 
			field => 
			{ 
					this.field = field
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
	   			editor: ["field-detail", ID]
	 	}
   	}]);
 }

}
