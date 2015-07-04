using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class GhostBehavior
    {
        public event EventHandler StateChanged;

        public Ghost.States State
        {
            get { return this.state; }
            set { this.state = value; }
        }

        public float Velocity
        {
            get { return this.ghostVelocity; }
        }

        public GhostBehavior()
        {
            Random rnd = new Random();

            this.state = Ghost.States.Waiting;
            this.ghostVelocity = Constants.DEFAULT_PLAYER_VELOCITY;
            this.transition = rnd.Next(10);
            this.Wait();
        }

        public void Wait()
        {
            this.nextState  = Ghost.States.Waiting;
            this.duration = 15.0f;
        }

        public void Fright()
        {
            this.nextState  = Ghost.States.Frightened;
            this.ghostVelocity = Constants.DEFAULT_PLAYER_VELOCITY * 0.8f;
            this.duration = 10.0;
        }

        public void Walk()
        {
            this.nextState = Ghost.States.Scatter;
            this.ghostVelocity = Constants.DEFAULT_PLAYER_VELOCITY;
            this.duration = 15.0;
        }

        public void Chase()
        {
            this.nextState = Ghost.States.Chase;
            this.ghostVelocity = Constants.DEFAULT_PLAYER_VELOCITY;
            this.duration = 30.0;
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
                
                this.OnStateChanged();
            }
            else
                transition -= elapsed;
        }

        private void OnStateChanged()
        {
            if (this.StateChanged != null)
                this.StateChanged.Invoke(this, null);
        }

        private Ghost.States state;
        private Ghost.States nextState;
        private double duration;
        private double transition;
        private float ghostVelocity;

        
    }
}
