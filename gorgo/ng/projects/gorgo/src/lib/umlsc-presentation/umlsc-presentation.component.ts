import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'


import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-umlsc-presentation',
	templateUrl: './umlsc-presentation.component.html',
	styleUrls: ['./umlsc-presentation.component.css']
})
export class UmlscPresentationComponent implements OnInit {

	umlsc: UmlscDB;



	PkgeltsViaUmlscsFieldName = "Umlscs"; // Label used to generates the table of Pkgelt that points to Umlsc via Umlscs
	PkgeltsViaUmlscsStructName = "Umlsc"; // Label used to generates the table of Pkgelt that points to Umlsc via Umlscs

	constructor(
		private umlscService: UmlscService,

		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.getUmlsc();


		// observable for changes in 
		this.umlscService.UmlscServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUmlsc()
					
				}
			}
		)
	}

  getUmlsc(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.umlscService.getUmlsc(id)
		.subscribe( 
			umlsc => 
			{ 
					this.umlsc = umlsc
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
	   			editor: ["umlsc-detail", ID]
	 	}
   	}]);
 }

}
