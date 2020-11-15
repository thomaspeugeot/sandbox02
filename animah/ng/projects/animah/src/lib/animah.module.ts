import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion'; 

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent} from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';


import { ActionsTableComponent } from './actions-table/actions-table.component'
import { ActionAdderComponent } from './action-adder/action-adder.component'
import { ActionDetailComponent } from './action-detail/action-detail.component'
import { ActionPresentationComponent } from './action-presentation/action-presentation.component'

import { AgentsTableComponent } from './agents-table/agents-table.component'
import { AgentAdderComponent } from './agent-adder/agent-adder.component'
import { AgentDetailComponent } from './agent-detail/agent-detail.component'
import { AgentPresentationComponent } from './agent-presentation/agent-presentation.component'

import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineAdderComponent } from './engine-adder/engine-adder.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'
import { EnginePresentationComponent } from './engine-presentation/engine-presentation.component'


@NgModule({
  declarations: [


	ActionsTableComponent,
	ActionAdderComponent,
	ActionDetailComponent,
	ActionPresentationComponent,

	AgentsTableComponent,
	AgentAdderComponent,
	AgentDetailComponent,
	AgentPresentationComponent,

	EnginesTableComponent,
	EngineAdderComponent,
	EngineDetailComponent,
	EnginePresentationComponent,


	SplitterComponent,
	SidebarComponent
  ],
  imports: [
	FormsModule,
	ReactiveFormsModule,
    CommonModule,
    RouterModule,

    AppRoutingModule,

    MatSliderModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
	MatToolbarModule,
	MatExpansionModule,
	MatListModule,

    AngularSplitModule,
  ],
  exports: [  
	  
	
	ActionsTableComponent,
	ActionAdderComponent,
	ActionDetailComponent,
	ActionPresentationComponent,

	AgentsTableComponent,
	AgentAdderComponent,
	AgentDetailComponent,
	AgentPresentationComponent,

	EnginesTableComponent,
	EngineAdderComponent,
	EngineDetailComponent,
	EnginePresentationComponent,


  	SplitterComponent,
  	SidebarComponent,

  ]
})
export class AnimahModule { }
