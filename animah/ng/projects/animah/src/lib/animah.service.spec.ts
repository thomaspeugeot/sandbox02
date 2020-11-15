import { TestBed } from '@angular/core/testing';

import { AnimahService } from './animah.service';

describe('AnimahService', () => {
  let service: AnimahService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AnimahService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
