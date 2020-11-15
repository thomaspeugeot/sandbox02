import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ActionSinglotonDB } from '../actionsingloton-db'
import { ActionSinglotonService } from '../actionsingloton.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-actionsingloton-presentation',
	templateUrl: './actionsingloton-presentation.component.html',
	styleUrls: ['./actionsingloton-presentation.component.css']
})
export class ActionSinglotonPresentationComponent implements OnInit {

	actionsingloton: ActionSinglotonDB;




	constructor(
		private actionsinglotonService: ActionSinglotonService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getActionSingloton();


		// observable for changes in 
		this.actionsinglotonService.ActionSinglotonServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getActionSingloton()
					
				}
			}
		)
	}

  getActionSingloton(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.actionsinglotonService.getActionSingloton(id)
		.subscribe( 
			actionsingloton => 
			{ 
					this.actionsingloton = actionsingloton
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
	   			editor: ["actionsingloton-detail", ID]
	 	}
   	}]);
 }

}
