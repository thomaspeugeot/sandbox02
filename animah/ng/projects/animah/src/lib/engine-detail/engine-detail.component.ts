// generated by genNgDetail.go
import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'



import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-engine-detail',
  templateUrl: './engine-detail.component.html',
  styleUrls: ['./engine-detail.component.css']
})
export class EngineDetailComponent implements OnInit {

	// the EngineDB of interest
	engine: EngineDB;

	constructor(
		private engineService: EngineService,

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
	this.getEngine();

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


  save(): void {
	const id = +this.route.snapshot.paramMap.get('id');




	this.engineService.updateEngine( this.engine )
    .subscribe(engine => {
		this.engineService.EngineServiceChanged.next("update")

    	console.log("engine saved")
    });
  }
}