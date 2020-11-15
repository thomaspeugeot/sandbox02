import { BrowserModule } from '@angular/platform-browser';
import { NgModule, ModuleWithProviders, SkipSelf, Optional  } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { HttpClientModule } from '@angular/common/http'; 

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

import { FormsModule } from '@angular/forms';

// for the tree
import { MatTreeModule } from '@angular/material/tree'
import { ReactiveFormsModule } from '@angular/forms';
import { MatNativeDateModule } from '@angular/material/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

// split
import { AngularSplitModule } from 'angular-split';

// devlopment in this module
import { SplitterComponent } from './splitter/splitter.component';
import { SidebarComponent } from './sidebar/sidebar.component'


import { MachinesTableComponent } from './machines-table/machines-table.component'
import { MachineAdderComponent } from './machine-adder/machine-adder.component'
import { MachineDetailComponent } from './machine-detail/machine-detail.component'
import { MachinePresentationComponent } from './machine-presentation/machine-presentation.component'

import { WashersTableComponent } from './washers-table/washers-table.component'
import { WasherAdderComponent } from './washer-adder/washer-adder.component'
import { WasherDetailComponent } from './washer-detail/washer-detail.component'
import { WasherPresentationComponent } from './washer-presentation/washer-presentation.component'


@NgModule({
  declarations: [
	SplitterComponent,
	

	MachinesTableComponent,
	MachineAdderComponent,
	MachineDetailComponent,
	MachinePresentationComponent,

	WashersTableComponent,
	WasherAdderComponent,
	WasherDetailComponent,
	WasherPresentationComponent,


    SidebarComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,

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

    MatTreeModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatListModule,

    FormsModule,

    AngularSplitModule.forRoot()
  ],
  providers: [],
  bootstrap: []
})
export class laundromatMatModule {

    constructor( @Optional() @SkipSelf() parentModule: laundromatMatModule,
                 @Optional() http: HttpClient) {
        if (parentModule) {
            throw new Error('laundromatMatModule is already loaded. Import in your base AppModule only.');
        }
        if (!http) {
            throw new Error('You need to import the HttpClientModule in your AppModule! \n' +
            'See also https://github.com/angular/angular/issues/20575');
        }
    }
}
