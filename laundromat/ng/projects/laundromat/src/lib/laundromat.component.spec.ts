import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LaundromatComponent } from './laundromat.component';

describe('LaundromatComponent', () => {
  let component: LaundromatComponent;
  let fixture: ComponentFixture<LaundromatComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LaundromatComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LaundromatComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
