import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { MachineAPI} from '../machine-api'
import { MachineDB} from '../machine-db'
import { MachineService} from '../machine.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-washer-presentation',
	templateUrl: './washer-presentation.component.html',
	styleUrls: ['./washer-presentation.component.css']
})
export class WasherPresentationComponent implements OnInit {

	washer: WasherDB;

	Machine = {} as MachineDB; // storing values of the field Machine of type Machine


	// generated by genEditableReadablePointerToStruct.go
	machines: MachineDB[];


	constructor(
		private washerService: WasherService,

		private machineService: MachineService,
		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getWasher();

    	this.getMachines();


		// observable for changes in 
		this.washerService.WasherServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getWasher()
					
    	this.getMachines();

				}
			}
		)
	}

  getWasher(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.washerService.getWasher(id)
		.subscribe( 
			washer => 
			{ 
					this.washer = washer
        	}
  	);
  }


	// generated by genEditableReadablePointerToStruct.go
	getMachines(): void {
		this.machineService.getMachines().subscribe(
			machines => {
				this.machines = machines;

				// init variable for each pointer
				this.machines.forEach(machine => {
					if (machine.ID == this.washer.MachineID) {
						this.Machine = machine
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
	   			editor: ["washer-detail", ID]
	 	}
   	}]);
 }

}
