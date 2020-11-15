import { TestBed } from '@angular/core/testing';

import { AnimahcontrolService } from './animahcontrol.service';

describe('AnimahcontrolService', () => {
  let service: AnimahcontrolService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AnimahcontrolService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
