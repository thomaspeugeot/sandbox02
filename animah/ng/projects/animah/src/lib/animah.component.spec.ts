import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AnimahComponent } from './animah.component';

describe('AnimahComponent', () => {
  let component: AnimahComponent;
  let fixture: ComponentFixture<AnimahComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AnimahComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnimahComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
