// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { EventDB } from '../event-db'
import { EventService } from '../event.service'


// generated table component
@Component({
  selector: 'app-events-table',
  templateUrl: './events-table.component.html',
  styleUrls: ['./events-table.component.css']
})
export class EventsTableComponent implements OnInit {

  // the data source for the table
  events: EventDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Event
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Duration', 'FireTime', 'Name', 'Edit', 'Delete'];

  constructor(
    private eventService: EventService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.eventService.EventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getEvents()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getEvents()
  }

  getEvents(): void {
    if (this.ID == null) {
      this.eventService.getEvents().subscribe(
        Events => {
          this.events = Events;
        }
      )
    }
  
  }

  // newEvent initiate a new event
  // create a new Event objet
  newEvent() {
  }

  deleteEvent(eventID: number, event: EventDB) {
    // la liste des events est amputée du event avant le delete afin
    // de mettre à jour l'IHM
    this.events = this.events.filter(h => h !== event);

    this.eventService.deleteEvent(eventID).subscribe();
  }

  editEvent(eventID: number, event: EventDB) {

  }

  // display event in router
  displayEventInRouter(eventID: number) {
    this.router.navigate( ["event-display", eventID])
  }

  // set editor outlet
  setEditorRouterOutlet(eventID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["event-detail", eventID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(eventID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["event-presentation", eventID]
      }
    }]);
  }
}
