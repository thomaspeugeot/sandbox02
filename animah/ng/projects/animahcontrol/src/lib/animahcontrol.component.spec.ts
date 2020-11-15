import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AnimahcontrolComponent } from './animahcontrol.component';

describe('AnimahcontrolComponent', () => {
  let component: AnimahcontrolComponent;
  let fixture: ComponentFixture<AnimahcontrolComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AnimahcontrolComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnimahcontrolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
