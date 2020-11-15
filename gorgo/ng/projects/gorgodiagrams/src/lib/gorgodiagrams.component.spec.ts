import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GorgodiagramsComponent } from './gorgodiagrams.component';

describe('GorgodiagramsComponent', () => {
  let component: GorgodiagramsComponent;
  let fixture: ComponentFixture<GorgodiagramsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GorgodiagramsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GorgodiagramsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
