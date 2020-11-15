import { TestBed } from '@angular/core/testing';

import { LaundromatService } from './laundromat.service';

describe('LaundromatService', () => {
  let service: LaundromatService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LaundromatService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
