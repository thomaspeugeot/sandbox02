import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GorgoactionDB } from '../gorgoaction-db'
import { GorgoactionService } from '../gorgoaction.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-gorgoaction-presentation',
	templateUrl: './gorgoaction-presentation.component.html',
	styleUrls: ['./gorgoaction-presentation.component.css']
})
export class GorgoactionPresentationComponent implements OnInit {

	gorgoaction: GorgoactionDB;




	constructor(
		private gorgoactionService: GorgoactionService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getGorgoaction();


		// observable for changes in 
		this.gorgoactionService.GorgoactionServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGorgoaction()
					
				}
			}
		)
	}

  getGorgoaction(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gorgoactionService.getGorgoaction(id)
		.subscribe( 
			gorgoaction => 
			{ 
					this.gorgoaction = gorgoaction
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
	   			editor: ["gorgoaction-detail", ID]
	 	}
   	}]);
 }

}
