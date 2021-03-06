// generated by genNgDetail.go
import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'



import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-umlsc-detail',
  templateUrl: './umlsc-detail.component.html',
  styleUrls: ['./umlsc-detail.component.css']
})
export class UmlscDetailComponent implements OnInit {

	// the UmlscDB of interest
	umlsc: UmlscDB;

	constructor(
		private umlscService: UmlscService,

		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
		return false;
		};
  }

  ngOnInit(): void {
	this.getUmlsc();

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


  save(): void {
	const id = +this.route.snapshot.paramMap.get('id');




	this.umlscService.updateUmlsc( this.umlsc )
    .subscribe(umlsc => {
		this.umlscService.UmlscServiceChanged.next("update")

    	console.log("umlsc saved")
    });
  }
}
