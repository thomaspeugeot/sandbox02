import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { StateDB } from '../state-db'
import { StateService } from '../state.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-state-presentation',
	templateUrl: './state-presentation.component.html',
	styleUrls: ['./state-presentation.component.css']
})
export class StatePresentationComponent implements OnInit {

	state: StateDB;



	UmlscsViaStatesFieldName = "States"; // Label used to generates the table of Umlsc that points to State via States
	UmlscsViaStatesStructName = "State"; // Label used to generates the table of Umlsc that points to State via States

	constructor(
		private stateService: StateService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getState();


		// observable for changes in 
		this.stateService.StateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getState()
					
				}
			}
		)
	}

  getState(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.stateService.getState(id)
		.subscribe( 
			state => 
			{ 
					this.state = state
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
	   			editor: ["state-detail", ID]
	 	}
   	}]);
 }

}
