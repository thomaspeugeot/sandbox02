import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EventDB } from '../event-db'
import { EventService } from '../event.service'


@Component({
  selector: 'app-event-adder',
  templateUrl: './event-adder.component.html',
  styleUrls: ['./event-adder.component.css']
})
export class EventAdderComponent implements OnInit {

	event = {} as EventDB;






  constructor(
    private eventService: EventService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.eventService.postEvent( this.event )
    .subscribe(event => {
		this.eventService.EventServiceChanged.next("post")
		
		this.event = {} // reset fields
	    console.log("event added")
    });
  }
}
