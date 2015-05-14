﻿using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class GhostBehavior
    {
        public Ghost.States State
        {
            get { return this.state; }
            set { this.state = value; }
        }

        public GhostBehavior()
        {
            this.Wait();
        }

        public void Wait()
        {
            this.nextState  = Ghost.States.Waiting;
            this.duration = 1.5;
        }

        public void Fright()
        {
            this.nextState  = Ghost.States.Frightened;
            this.duration = 2.5;
        }

        public void Walk()
        {
            this.nextState = Ghost.States.Scatter;
            this.duration = 3.5;
        }

        public void Chase()
        {
            this.nextState = Ghost.States.Chase;
            this.duration = 2.0;
        }

        public void Update(GameTime gameTime)
        {
            var elapsed = gameTime.ElapsedGameTime.TotalSeconds;

            //Devemos aguardar um certo tempo
            if (transition <= 0)
            {
                if (this.state == Ghost.States.Frightened || this.state == Ghost.States.Waiting || this.state == Ghost.States.Chase)
                {
                    this.Walk();
                }
                else if (this.state == Ghost.States.Scatter)
                {
                    this.Chase();
                }
                
                this.transition = this.duration;
                this.state = this.nextState;
            }
            else
                transition -= elapsed;
        }

        private Ghost.States state;
        private Ghost.States nextState;
        private double duration;
        private double transition;

        
    }
}
