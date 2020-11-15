import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GorgoComponent } from './gorgo.component';

describe('GorgoComponent', () => {
  let component: GorgoComponent;
  let fixture: ComponentFixture<GorgoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GorgoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GorgoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
