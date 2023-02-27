import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ByIDComponent } from './by-id.component';

describe('ByIDComponent', () => {
  let component: ByIDComponent;
  let fixture: ComponentFixture<ByIDComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ByIDComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ByIDComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
