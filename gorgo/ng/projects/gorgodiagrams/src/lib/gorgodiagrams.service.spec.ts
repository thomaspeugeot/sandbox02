import { TestBed } from '@angular/core/testing';

import { GorgodiagramsService } from './gorgodiagrams.service';

describe('GorgodiagramsService', () => {
  let service: GorgodiagramsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GorgodiagramsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
