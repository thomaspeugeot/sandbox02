// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { MachineAPI} from '../machine-api'
import { MachineDB} from '../machine-db'
import { MachineService} from '../machine.service'


// generated table component
@Component({
  selector: 'app-washers-table',
  templateUrl: './washers-table.component.html',
  styleUrls: ['./washers-table.component.css']
})
export class WashersTableComponent implements OnInit {

  // the data source for the table
  washers: WasherDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Washer
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'CleanedLaundryWeight', 'LaundryWeight', 'Machine', 'Name', 'State', 'Edit', 'Delete'];

  constructor(
    private washerService: WasherService,

		private machineService: MachineService,
    private router: Router,
  ) {
    // observable for changes in structs
    this.washerService.WasherServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getWashers()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getWashers()
  }

  getWashers(): void {
    if (this.ID == null) {
      this.washerService.getWashers().subscribe(
        Washers => {
          this.washers = Washers;
        }
      )
    }
  
		if (this.ID != 0 && this.field == "Machine" && this.struct == "Machine") {
			this.machineService.getMachineWashersViaMachine(this.ID).subscribe(
				washers => {
					this.washers = washers;
				}
			)
		}
  }

  // newWasher initiate a new washer
  // create a new Washer objet
  newWasher() {
  }

  deleteWasher(washerID: number, washer: WasherDB) {
    // la liste des washers est amputée du washer avant le delete afin
    // de mettre à jour l'IHM
    this.washers = this.washers.filter(h => h !== washer);

    this.washerService.deleteWasher(washerID).subscribe();
  }

  editWasher(washerID: number, washer: WasherDB) {

  }

  // display washer in router
  displayWasherInRouter(washerID: number) {
    this.router.navigate( ["washer-display", washerID])
  }

  // set editor outlet
  setEditorRouterOutlet(washerID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["washer-detail", washerID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(washerID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["washer-presentation", washerID]
      }
    }]);
  }
}
