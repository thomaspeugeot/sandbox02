import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-engine-presentation',
	templateUrl: './engine-presentation.component.html',
	styleUrls: ['./engine-presentation.component.css']
})
export class EnginePresentationComponent implements OnInit {

	engine: EngineDB;



	AgentsViaEngineFieldName = "Engine"; // Label used to generates the table of Agent that points to Engine via Engine
	AgentsViaEngineStructName = "Engine"; // Label used to generates the table of Agent that points to Engine via Engine

	constructor(
		private engineService: EngineService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getEngine();


		// observable for changes in 
		this.engineService.EngineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEngine()
					
				}
			}
		)
	}

  getEngine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.engineService.getEngine(id)
		.subscribe( 
			engine => 
			{ 
					this.engine = engine
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
	   			editor: ["engine-detail", ID]
	 	}
   	}]);
 }

}
