import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

// mandatory
import { HttpClientModule } from '@angular/common/http';

import { AnimahcontrolModule } from 'animahcontrol'
import { AnimahModule } from 'animah'

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    HttpClientModule,

    AnimahcontrolModule,
    AnimahModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
