// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { PositionAPI} from '../position-api'
import { PositionDB} from '../position-db'
import { PositionService} from '../position.service'


// generated table component
@Component({
  selector: 'app-classshapes-table',
  templateUrl: './classshapes-table.component.html',
  styleUrls: ['./classshapes-table.component.css']
})
export class ClassshapesTableComponent implements OnInit {

  // the data source for the table
  classshapes: ClassshapeDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Classshape
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Heigth', 'Name', 'Position', 'Structname', 'Width', 'Edit', 'Delete'];

  constructor(
    private classshapeService: ClassshapeService,

		private positionService: PositionService,
    private router: Router,
  ) {
    // observable for changes in structs
    this.classshapeService.ClassshapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getClassshapes()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getClassshapes()
  }

  getClassshapes(): void {
    if (this.ID == null) {
      this.classshapeService.getClassshapes().subscribe(
        Classshapes => {
          this.classshapes = Classshapes;
        }
      )
    }
  
		if (this.ID != 0 && this.field == "Position" && this.struct == "Position") {
			this.positionService.getPositionClassshapesViaPosition(this.ID).subscribe(
				classshapes => {
					this.classshapes = classshapes;
				}
			)
		}
  }

  // newClassshape initiate a new classshape
  // create a new Classshape objet
  newClassshape() {
  }

  deleteClassshape(classshapeID: number, classshape: ClassshapeDB) {
    // la liste des classshapes est amputée du classshape avant le delete afin
    // de mettre à jour l'IHM
    this.classshapes = this.classshapes.filter(h => h !== classshape);

    this.classshapeService.deleteClassshape(classshapeID).subscribe();
  }

  editClassshape(classshapeID: number, classshape: ClassshapeDB) {

  }

  // display classshape in router
  displayClassshapeInRouter(classshapeID: number) {
    this.router.navigate( ["classshape-display", classshapeID])
  }

  // set editor outlet
  setEditorRouterOutlet(classshapeID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["classshape-detail", classshapeID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(classshapeID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["classshape-presentation", classshapeID]
      }
    }]);
  }
}