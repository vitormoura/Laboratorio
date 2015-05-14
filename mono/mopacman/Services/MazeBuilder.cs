using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Content;
using Microsoft.Xna.Framework.Graphics;
using mopacman.Components;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

namespace mopacman.Services
{
    class MazeBuilder
    {
        public static Maze GetDefaultFor(ContentManager cm)
        {
            Int32 width, height;
            String blueprint = Path.Combine(Environment.CurrentDirectory, cm.RootDirectory, "maze_blueprint.txt");
            String[] lines = File.ReadAllLines(blueprint);
            Dictionary<Char, MazeSection> linkedSections = new Dictionary<char, MazeSection>();

            if (lines != null && lines.Length > 0)
            {
                height = lines.Length;
                width = lines[0].Length;

                MazeSection[,] sections = new MazeSection[height,width];
                
                for (int y = 0; y < height; y++)
                {
                    String line = lines[y];
                    MazeSection lastW = null;

                    for (int x = 0; x < line.Length; x++)
                    {
                        Char c = line[x];
                        MazeSection s = new MazeSection(x, y);
                        s.W = lastW;
                        s.Checkpoint = c == 'R';
                        s.HasCookie = c == '-';
                        s.Allowed = (c == '-' || c == '_' || char.IsDigit(c));

                        //Redefinindo o 'oeste' o último leste
                        if (lastW != null)
                        {
                            lastW.E = s;
                        }

                        //Redefinindo o 'sul' do último norte e avançando o norte para o elemento ao lado
                        if (y > 0)
                        {
                            s.N = sections[y - 1, x];
                            s.N.S = s;
                        }

                        lastW = s;
                        sections[y,x] = s;

                        //Estamos lidando com uma linked section ?
                        if (char.IsDigit(c))
                        {
                            if (linkedSections.ContainsKey(c))
                            {
                                MazeSection linkedS = linkedSections[c];

                                if (s.E == null)
                                {
                                    s.E = linkedS;
                                    linkedS.W = s;
                                }
                                else if (linkedS.W == null)
                                {
                                    s.W = linkedS;
                                    linkedS.E = s;
                                }
                                else if (linkedS.N == null)
                                {
                                    s.N = linkedS;
                                    linkedS.S = s;
                                }
                                else if (linkedS.S == null)
                                {
                                    s.S = linkedS;
                                    linkedS.N = s;
                                }
                            }
                            else
                            {
                                linkedSections.Add(c,s);
                            }
                        }
                    }
                }

                return new Maze(sections);
            }
            else
                throw new FileLoadException("Formato de arquivo de bluprint inválido");
        }
    }
}
