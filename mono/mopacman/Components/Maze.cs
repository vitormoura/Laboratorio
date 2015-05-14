using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Maze : IEnumerable<MazeSection>
    {
        public Int32 Width { get; private set; }

        public Int32 Height { get; private set; }

        public Int32 Size 
        {
            get 
            { 
                return this.Width * this.Height; 
            } 
        }

        public IList<MazeSection> Checkpoints
        {
            get { return this.checkpoints; }
        }

        public MazeSection this [int x, int y] 
        {
            get 
            {
                return this.sections[x, y];
            }
        }

        public Maze(MazeSection[,] sections)
        {
            this.sections = sections;
            this.Width = sections.GetLength(0);
            this.Height = sections.GetLength(1);

            this.checkpoints = this.Where(x => x.Checkpoint).ToList();
        }

        public MazeSection GetStartSection()
        {
            return this[29, 14];
        }

        public MazeSection GetGhostLairSection()
        {
            return this[13, 13];
        }
                
        private MazeSection[,]      sections;
        private List<MazeSection>   checkpoints;

        #region implementação de IEnumerable

        public IEnumerator<MazeSection> GetEnumerator()
        {
            for (int i = 0; i < this.Height; i++)
            {
                for (int j = 0; j < this.Width; j++)
                {
                    yield return this.sections[j, i];
                }
            }
        }

        System.Collections.IEnumerator System.Collections.IEnumerable.GetEnumerator()
        {
            return this.GetEnumerator();
        }

        #endregion
    }
}
