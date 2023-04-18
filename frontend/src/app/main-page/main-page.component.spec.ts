import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MainPageComponent } from './main-page.component';
import { FormsModule } from '@angular/forms';

describe('MainPageComponent', () => {
  let component: MainPageComponent;
  let fixture: ComponentFixture<MainPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MainPageComponent ],
      imports: [HttpClientTestingModule, FormsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MainPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  /*
  it('posting a thread with the built in form should work', () => {
    const mainPage = fixture.debugElement.injector.get(MainPageComponent);
    component.username = "testing";
    component.title = "testing";
    component.body = "this is a testing post";
    mainPage.addThread();
    fixture.detectChanges();
    expect(fixture.debugElement.nativeElement.innerHTML).toContain("this is a testing post")
  });*/
});
