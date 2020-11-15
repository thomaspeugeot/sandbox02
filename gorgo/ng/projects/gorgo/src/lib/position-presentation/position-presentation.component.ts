import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-position-presentation',
	templateUrl: './position-presentation.component.html',
	styleUrls: ['./position-presentation.component.css']
})
export class PositionPresentationComponent implements OnInit {

	position: PositionDB;



	ClassshapesViaPositionFieldName = "Position"; // Label used to generates the table of Classshape that points to Position via Position
	ClassshapesViaPositionStructName = "Position"; // Label used to generates the table of Classshape that points to Position via Position

	constructor(
		private positionService: PositionService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getPosition();


		// observable for changes in 
		this.positionService.PositionServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPosition()
					
				}
			}
		)
	}

  getPosition(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.positionService.getPosition(id)
		.subscribe( 
			position => 
			{ 
					this.position = position
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
	   			editor: ["position-detail", ID]
	 	}
   	}]);
 }

}
