import { TestBed } from '@angular/core/testing';

import { GorgoService } from './gorgo.service';

describe('GorgoService', () => {
  let service: GorgoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GorgoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
