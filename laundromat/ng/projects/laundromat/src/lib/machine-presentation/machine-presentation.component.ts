import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-machine-presentation',
	templateUrl: './machine-presentation.component.html',
	styleUrls: ['./machine-presentation.component.css']
})
export class MachinePresentationComponent implements OnInit {

	machine: MachineDB;



	WashersViaMachineFieldName = "Machine"; // Label used to generates the table of Washer that points to Machine via Machine
	WashersViaMachineStructName = "Machine"; // Label used to generates the table of Washer that points to Machine via Machine

	constructor(
		private machineService: MachineService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getMachine();


		// observable for changes in 
		this.machineService.MachineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMachine()
					
				}
			}
		)
	}

  getMachine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.machineService.getMachine(id)
		.subscribe( 
			machine => 
			{ 
					this.machine = machine
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
	   			editor: ["machine-detail", ID]
	 	}
   	}]);
 }

}
