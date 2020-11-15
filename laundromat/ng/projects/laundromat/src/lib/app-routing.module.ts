import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


import { MachinesTableComponent } from './machines-table/machines-table.component'
import { MachineAdderComponent } from './machine-adder/machine-adder.component'
import { MachineDetailComponent } from './machine-detail/machine-detail.component'
import { MachinePresentationComponent } from './machine-presentation/machine-presentation.component'

import { WashersTableComponent } from './washers-table/washers-table.component'
import { WasherAdderComponent } from './washer-adder/washer-adder.component'
import { WasherDetailComponent } from './washer-detail/washer-detail.component'
import { WasherPresentationComponent } from './washer-presentation/washer-presentation.component'

const routes: Routes = [
	{ path: 'machines', component: MachinesTableComponent, outlet: 'table' },
	{ path: 'machine-adder', component: MachineAdderComponent, outlet: 'editor' },
	{ path: 'machine-detail/:id', component: MachineDetailComponent, outlet: 'editor' },
	{ path: 'machine-presentation/:id', component: MachinePresentationComponent, outlet: 'presentation' },
	{ path: 'machine-presentation-special/:id', component: MachinePresentationComponent, outlet: 'machinepres' },

	{ path: 'washers', component: WashersTableComponent, outlet: 'table' },
	{ path: 'washer-adder', component: WasherAdderComponent, outlet: 'editor' },
	{ path: 'washer-detail/:id', component: WasherDetailComponent, outlet: 'editor' },
	{ path: 'washer-presentation/:id', component: WasherPresentationComponent, outlet: 'presentation' },
	{ path: 'washer-presentation-special/:id', component: WasherPresentationComponent, outlet: 'washerpres' },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
