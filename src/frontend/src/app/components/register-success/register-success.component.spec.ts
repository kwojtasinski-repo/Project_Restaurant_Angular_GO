import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterSuccessComponent } from './register-success.component';
import { By } from '@angular/platform-browser';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('RegisterSuccessComponent', () => {
  let component: RegisterSuccessComponent;
  let fixture: ComponentFixture<RegisterSuccessComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RegisterSuccessComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterSuccessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create content', () => {
    const registerSuccessForm = (fixture.nativeElement as HTMLElement).querySelector('.border-register-success-form');

    const registerText = registerSuccessForm?.querySelector('div > h3');
    const linkNavigateBack = registerSuccessForm?.querySelector('div > a');

    expect(registerText).not.toBeUndefined();
    expect(registerText).not.toBeNull();
    expect(registerText?.innerHTML).not.toBeUndefined();
    expect(registerText?.innerHTML).not.toBeNull();
    expect(registerText?.innerHTML.length).toBeGreaterThan(0);
    expect(linkNavigateBack?.innerHTML).not.toBeUndefined();
    expect(linkNavigateBack?.innerHTML).not.toBeNull();
    expect(linkNavigateBack?.innerHTML).not.toBeUndefined();
    expect(linkNavigateBack?.innerHTML).not.toBeNull();
    expect(linkNavigateBack?.innerHTML.length).toBeGreaterThan(0);
  });

  it('should create link with proper href url', async () => {
    const anchorTag = fixture.debugElement.query(By.css('a')).nativeElement;
    
    expect(anchorTag).not.toBeUndefined();
    expect(anchorTag).not.toBeNull();
    const hrefAttribute = anchorTag.getAttribute('href');
    expect(hrefAttribute).not.toBeUndefined();
    expect(hrefAttribute).not.toBeNull();
    expect(hrefAttribute).toEqual('/login');
  });
});
