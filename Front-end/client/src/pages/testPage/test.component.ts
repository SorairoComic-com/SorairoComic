import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Book } from './model';

@Component({
  selector: 'home-page',
  templateUrl: './test.component.html',

  // satu paket yang dibawah
  standalone: true,
  imports: [CommonModule],
})
export class TestComponent {
  myColorExpression: string = 'blue';

  isError: boolean = true;

  textColor: string = 'green';
  fontSize: number = 18;


  arrData: Book[] = [
    {
      name: 'Book 1',
    },
    {
      name: 'Book 2',
    },
  ];

  active: boolean = true;
  handleClick() {
    this.active = !this.active;
  }

    inputUser1(event: any) {
      console.log(event.target.value);
    }

    inputUser2(event: any) {
      console.log(event.target.value);
    }
  
  
}
