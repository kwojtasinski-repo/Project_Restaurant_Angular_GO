import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenuComponent } from './menu.component';
import { initialState } from 'src/app/stores/login/login.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { HttpClientModule } from '@angular/common/http';

describe('MenuComponent', () => {
  let component: MenuComponent;
  let fixture: ComponentFixture<MenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        MenuComponent, 
        SearchBarComponent, 
        MoneyPipe 
      ],
      imports: [
        RouterTestingModule,
        FormsModule,
        HttpClientModule
      ],
      providers: [
        provideMockStore({ initialState }),
        {
          provide: "API_URL", useValue: ''
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
