import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ActionDB } from '../action-db'
import { ActionService } from '../action.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-action-presentation',
	templateUrl: './action-presentation.component.html',
	styleUrls: ['./action-presentation.component.css']
})
export class ActionPresentationComponent implements OnInit {

	action: ActionDB;




	constructor(
		private actionService: ActionService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getAction();


		// observable for changes in 
		this.actionService.ActionServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAction()
					
				}
			}
		)
	}

  getAction(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.actionService.getAction(id)
		.subscribe( 
			action => 
			{ 
					this.action = action
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
	   			editor: ["action-detail", ID]
	 	}
   	}]);
 }

}
