// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'


// generated table component
@Component({
  selector: 'app-machines-table',
  templateUrl: './machines-table.component.html',
  styleUrls: ['./machines-table.component.css']
})
export class MachinesTableComponent implements OnInit {

  // the data source for the table
  machines: MachineDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Machine
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Cleanedlaundry', 'DrumLoad', 'Name', 'RemainingTime', 'RemainingTimeMinutes', 'State', 'Edit', 'Delete'];

  constructor(
    private machineService: MachineService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.machineService.MachineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getMachines()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getMachines()
  }

  getMachines(): void {
    if (this.ID == null) {
      this.machineService.getMachines().subscribe(
        Machines => {
          this.machines = Machines;
        }
      )
    }
  
  }

  // newMachine initiate a new machine
  // create a new Machine objet
  newMachine() {
  }

  deleteMachine(machineID: number, machine: MachineDB) {
    // la liste des machines est amputée du machine avant le delete afin
    // de mettre à jour l'IHM
    this.machines = this.machines.filter(h => h !== machine);

    this.machineService.deleteMachine(machineID).subscribe();
  }

  editMachine(machineID: number, machine: MachineDB) {

  }

  // display machine in router
  displayMachineInRouter(machineID: number) {
    this.router.navigate( ["machine-display", machineID])
  }

  // set editor outlet
  setEditorRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["machine-detail", machineID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(machineID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["machine-presentation", machineID]
      }
    }]);
  }
}
