import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { AgentDB } from '../agent-db'
import { AgentService } from '../agent.service'

import { EngineAPI} from '../engine-api'
import { EngineDB} from '../engine-db'
import { EngineService} from '../engine.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-agent-presentation',
	templateUrl: './agent-presentation.component.html',
	styleUrls: ['./agent-presentation.component.css']
})
export class AgentPresentationComponent implements OnInit {

	agent: AgentDB;

	Engine = {} as EngineDB; // storing values of the field Engine of type Engine


	// generated by genEditableReadablePointerToStruct.go
	engines: EngineDB[];


	constructor(
		private agentService: AgentService,

		private engineService: EngineService,
		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getAgent();

    	this.getEngines();


		// observable for changes in 
		this.agentService.AgentServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAgent()
					
    	this.getEngines();

				}
			}
		)
	}

  getAgent(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.agentService.getAgent(id)
		.subscribe( 
			agent => 
			{ 
					this.agent = agent
        	}
  	);
  }


	// generated by genEditableReadablePointerToStruct.go
	getEngines(): void {
		this.engineService.getEngines().subscribe(
			engines => {
				this.engines = engines;

				// init variable for each pointer
				this.engines.forEach(engine => {
					if (engine.ID == this.agent.EngineID) {
						this.Engine = engine
					}
				});
      		}
    	)
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
	   			editor: ["agent-detail", ID]
	 	}
   	}]);
 }

}
